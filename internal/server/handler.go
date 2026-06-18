package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"interviewer-agent/internal/agent"
	"interviewer-agent/internal/extract"
	"interviewer-agent/internal/llm"
	"interviewer-agent/internal/model"
)

var _ = extract.PDF // ensure package is imported

// Handler wires all HTTP routes.
type Handler struct {
	store  *agent.Store
	webFS  fs.FS
	mux    *http.ServeMux
}

// New creates the Handler and registers all routes.
func New(store *agent.Store, webFS fs.FS) *Handler {
	h := &Handler{store: store, webFS: webFS}
	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("/api/companies", h.listCompanies)
	mux.HandleFunc("/api/providers", h.listProviders)
	mux.HandleFunc("/api/extract", h.extractText)
	mux.HandleFunc("/api/llm/stream", h.llmStream) // stateless streaming for coding mode
	mux.HandleFunc("/api/sessions", h.sessionsHandler) // GET list / POST create
	mux.HandleFunc("/api/sessions/", h.sessionRouter)

	// Static frontend
	mux.Handle("/", http.FileServer(http.FS(webFS)))

	h.mux = mux
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	corsMiddleware(h.mux).ServeHTTP(w, r)
}

// POST /api/llm/stream — stateless LLM streaming endpoint.
// The client manages conversation history; this is a thin pass-through.
// Body: { system, messages:[{role,content}], provider, provider_key, provider_model, provider_base_url }
// Response: SSE stream (same data: chunk\n\n format)
func (h *Handler) llmStream(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		System          string `json:"system"`
		Messages        []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"messages"`
		Provider        string `json:"provider"`
		ProviderKey     string `json:"provider_key"`
		ProviderModel   string `json:"provider_model"`
		ProviderBaseURL string `json:"provider_base_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Resolve provider
	var provider llm.Provider
	if req.ProviderKey != "" && req.Provider != "" {
		provider = llm.NewProviderFromKey(req.Provider, req.ProviderKey, req.ProviderModel, req.ProviderBaseURL)
	} else {
		registry := h.store.Providers()
		p, ok := registry[req.Provider]
		if !ok {
			// try first available
			for _, p2 := range registry {
				p = p2
				break
			}
		}
		if p == nil {
			http.Error(w, "no provider available", http.StatusBadRequest)
			return
		}
		provider = p
	}

	// Build model.Message slice
	var msgs []model.Message
	for _, m := range req.Messages {
		role := model.RoleUser
		if m.Role == "assistant" {
			role = model.RoleAssistant
		}
		msgs = append(msgs, model.Message{Role: role, Content: m.Content})
	}

	// Fake session (no files, just text)
	sess := &model.Session{}

	sseHeaders(w)
	_, err := provider.Stream(context.Background(), req.System, sess, msgs, w)
	if err != nil {
		fmt.Fprintf(w, "data: [ERROR] %s\n\n", err.Error())
	}
	fmt.Fprintf(w, "data: [DONE]\n\n")
	flush(w)
}

// POST /api/extract — extract text from PDF (Go library) or image (vision LLM)
func (h *Handler) extractText(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Type            string `json:"type"`      // "pdf" | "image"
		Data            string `json:"data"`      // base64
		MimeType        string `json:"mime_type"`
		Provider        string `json:"provider"`         // vision provider id (for images)
		ProviderKey     string `json:"provider_key"`     // user-supplied key
		ProviderModel   string `json:"provider_model"`
		ProviderBaseURL string `json:"provider_base_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Data == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	var text string
	var err error

	switch req.Type {
	case "pdf":
		text, err = extract.PDF(req.Data)

	case "image":
		vp, vpErr := h.pickVisionProvider(req.Provider, req.ProviderKey, req.ProviderModel, req.ProviderBaseURL)
		if vpErr != nil {
			http.Error(w, vpErr.Error(), http.StatusBadRequest)
			return
		}
		text, err = vp.ExtractImageText(context.Background(), req.Data, req.MimeType)

	default:
		http.Error(w, "type must be 'pdf' or 'image'", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"text": text})
}

// pickVisionProvider returns a vision-capable provider.
// Priority: user-supplied key (creates temp provider) > server-configured vision provider.
func (h *Handler) pickVisionProvider(id, key, model, baseURL string) (llm.Provider, error) {
	if key != "" && id != "" {
		return llm.NewProviderFromKey(id, key, model, baseURL), nil
	}
	registry := h.store.Providers()
	for _, pid := range []string{"anthropic", "openai", "qwen"} {
		if p, ok := registry[pid]; ok && p.Info().SupportsImg {
			return p, nil
		}
	}
	return nil, fmt.Errorf("没有可用的视觉 AI 供应商。请在 ⚙️ 设置中配置 Claude、GPT-4o 或 Qwen 的 API Key 以支持图片提取")
}

// GET /api/providers — returns ALL known providers, marking server-configured ones
func (h *Handler) listProviders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	writeJSON(w, llm.KnownProviders(h.store.Providers()))
}

// GET /api/companies
func (h *Handler) listCompanies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	companies := h.store.Companies()
	type item struct {
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Rounds      []int  `json:"rounds"`
	}
	var list []item
	for _, p := range companies {
		rounds := make([]int, 0, len(p.Rounds))
		for k := range p.Rounds {
			rounds = append(rounds, k)
		}
		list = append(list, item{Name: p.Name, DisplayName: p.DisplayName, Rounds: rounds})
	}
	writeJSON(w, list)
}

// GET /api/sessions → list  |  POST /api/sessions → create
func (h *Handler) sessionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listSessions(w, r)
	case http.MethodPost:
		h.createSession(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// GET /api/sessions — returns summary list (no image data to keep payload small)
func (h *Handler) listSessions(w http.ResponseWriter, r *http.Request) {
	sessions := h.store.ListSessions()
	type summary struct {
		ID           string `json:"id"`
		Company      string `json:"company"`
		Round        int    `json:"round"`
		Status       string `json:"status"`
		MessageCount int    `json:"message_count"`
		Preview      string `json:"preview"` // first visible assistant message
		CreatedAt    string `json:"created_at"`
	}
	list := make([]summary, 0, len(sessions))
	for _, sess := range sessions {
		preview := ""
		count := 0
		for _, m := range sess.Messages {
			if m.IsHidden {
				continue
			}
			count++
			if preview == "" && m.Role == model.RoleAssistant {
				r := []rune(m.Content)
				if len(r) > 80 {
					r = r[:80]
				}
				preview = string(r) + "…"
			}
		}
		list = append(list, summary{
			ID:           sess.ID,
			Company:      sess.Company,
			Round:        sess.Round,
			Status:       string(sess.Status),
			MessageCount: count,
			Preview:      preview,
			CreatedAt:    sess.CreatedAt.Format("2006-01-02 15:04"),
		})
	}
	writeJSON(w, list)
}

// POST /api/sessions
func (h *Handler) createSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Company         string               `json:"company"`
		Provider        string               `json:"provider"`
		ProviderKey     string               `json:"provider_key"`
		ProviderModel   string               `json:"provider_model"`
		ProviderBaseURL string               `json:"provider_base_url"`
		Round           int                  `json:"round"`
		JD              string               `json:"jd"`
		Resume          string               `json:"resume"`
		ResumePDF       *model.AttachedFile  `json:"resume_pdf"`
		ResumeImages    []model.AttachedFile `json:"resume_images"`
		JDImages        []model.AttachedFile `json:"jd_images"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if req.Company == "" || req.Provider == "" || req.Round < 1 {
		http.Error(w, "company, provider and round are required", http.StatusBadRequest)
		return
	}
	hasJD := req.JD != "" || len(req.JDImages) > 0
	hasResume := req.Resume != "" ||
		(req.ResumePDF != nil && req.ResumePDF.Data != "") ||
		len(req.ResumeImages) > 0
	if !hasJD || !hasResume {
		http.Error(w, "jd (text or image) and resume (text, pdf, or image) are required", http.StatusBadRequest)
		return
	}
	sess, err := h.store.CreateSession(req.Company, req.Provider, req.Round, req.JD, req.Resume, req.ResumePDF, req.ResumeImages, req.JDImages, req.ProviderKey, req.ProviderModel, req.ProviderBaseURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, sess)
}

// Routes under /api/sessions/{id}/...
func (h *Handler) sessionRouter(w http.ResponseWriter, r *http.Request) {
	// path: /api/sessions/{id} or /api/sessions/{id}/chat etc.
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/sessions/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "missing session id", http.StatusBadRequest)
		return
	}
	id := parts[0]
	action := ""
	if len(parts) > 1 {
		action = parts[1]
	}

	sess, err := h.store.GetSession(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	switch action {
	case "":
		// GET /api/sessions/{id}
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		writeJSON(w, sess)

	case "start":
		// POST /api/sessions/{id}/start – generate opening question
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h.sseStart(w, r, sess)

	case "chat":
		// POST /api/sessions/{id}/chat
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h.sseChat(w, r, sess)

	case "evaluate":
		// POST /api/sessions/{id}/evaluate
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h.sseEvaluate(w, r, sess)

	default:
		http.Error(w, "not found", http.StatusNotFound)
	}
}

func (h *Handler) sseStart(w http.ResponseWriter, r *http.Request, sess *model.Session) {
	sseHeaders(w)
	if err := h.store.StartInterview(sess, w); err != nil {
		fmt.Fprintf(w, "data: [ERROR] %s\n\n", err.Error())
	}
	fmt.Fprintf(w, "data: [DONE]\n\n")
	flush(w)
}

func (h *Handler) sseChat(w http.ResponseWriter, r *http.Request, sess *model.Session) {
	var req struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Message == "" {
		http.Error(w, "message is required", http.StatusBadRequest)
		return
	}
	sseHeaders(w)
	if err := h.store.Chat(sess, req.Message, w); err != nil {
		fmt.Fprintf(w, "data: [ERROR] %s\n\n", err.Error())
	}
	fmt.Fprintf(w, "data: [DONE]\n\n")
	flush(w)
}

func (h *Handler) sseEvaluate(w http.ResponseWriter, r *http.Request, sess *model.Session) {
	sseHeaders(w)
	if err := h.store.Evaluate(sess, w); err != nil {
		fmt.Fprintf(w, "data: [ERROR] %s\n\n", err.Error())
	}
	fmt.Fprintf(w, "data: [DONE]\n\n")
	flush(w)
}

func sseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
}

func flush(w http.ResponseWriter) {
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
