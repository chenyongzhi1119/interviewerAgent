package llm

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"interviewer-agent/internal/model"
)

// OpenAICompatProvider works with any OpenAI-format Chat Completions API,
// including OpenAI, DeepSeek, GLM (ZhipuAI), and Qwen (Alibaba).
type OpenAICompatProvider struct {
	baseURL string
	apiKey  string
	model   string
	info    ProviderInfo
	client  *http.Client
}

// Well-known provider presets
var openAICompatPresets = map[string]struct {
	baseURL     string
	defaultModel string
	displayName string
	supportsImg  bool
}{
	"openai":   {"https://api.openai.com/v1", "gpt-4.1", "GPT-4.1（OpenAI）", true},
	"deepseek": {"https://api.deepseek.com/v1", "deepseek-v4-flash", "DeepSeek", false},   // deepseek-chat 不支持图片
	"glm":      {"https://open.bigmodel.cn/api/paas/v4", "glm-4-flash", "GLM-4（智谱 AI）", false}, // glm-4-flash 不支持图片；视觉需用 glm-4v
	"qwen":     {"https://dashscope.aliyuncs.com/compatible-mode/v1", "qwen-plus", "Qwen（阿里云）", true},
}

// NewOpenAICompat creates an OpenAI-compatible provider.
// id must be one of the well-known presets or a custom id with explicit baseURL and model.
func NewOpenAICompat(id, apiKey, baseURL, modelID string) *OpenAICompatProvider {
	preset, ok := openAICompatPresets[id]
	if ok {
		if baseURL == "" {
			baseURL = preset.baseURL
		}
		if modelID == "" {
			modelID = preset.defaultModel
		}
	}
	supportsImg := ok && preset.supportsImg
	displayName := id
	if ok {
		displayName = preset.displayName
	}

	return &OpenAICompatProvider{
		baseURL: strings.TrimRight(baseURL, "/"),
		apiKey:  apiKey,
		model:   modelID,
		client:  &http.Client{Timeout: 120 * time.Second},
		info: ProviderInfo{
			ID:          id,
			DisplayName: displayName,
			Model:       modelID,
			SupportsPDF: false, // PDF not natively supported
			SupportsImg: supportsImg,
		},
	}
}

func (p *OpenAICompatProvider) Info() ProviderInfo { return p.info }

func (p *OpenAICompatProvider) ExtractImageText(ctx context.Context, base64Data, mimeType string) (string, error) {
	if !p.info.SupportsImg {
		return "", ErrNoVision
	}
	if ctx == nil {
		ctx = context.Background()
	}
	body := map[string]any{
		"model": p.model,
		"messages": []map[string]any{{
			"role": "user",
			"content": []map[string]any{
				{"type": "image_url", "image_url": map[string]string{
					"url": "data:" + mimeType + ";base64," + base64Data,
				}},
				{"type": "text", "text": "请提取图片中的所有文字内容，原样输出，保留原有段落格式，不要添加任何解释或额外内容。"},
			},
		}},
		"stream": false,
	}
	bodyBytes, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/chat/completions", bytes.NewReader(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.apiKey)

	resp, err := p.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s %s", resp.Status, string(data))
	}
	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(data, &result); err != nil || len(result.Choices) == 0 {
		return "", fmt.Errorf("unexpected response")
	}
	return result.Choices[0].Message.Content, nil
}

func (p *OpenAICompatProvider) Stream(ctx context.Context, systemPrompt string, sess *model.Session, history []model.Message, w io.Writer) (string, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	msgs := p.buildMessages(systemPrompt, sess, history)
	body := map[string]any{
		"model":    p.model,
		"stream":   true,
		"messages": msgs,
	}
	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/chat/completions", bytes.NewReader(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.apiKey)

	resp, err := p.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("%s %s", resp.Status, string(errBody))
	}

	var full string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data: ") {
			continue
		}
		data := strings.TrimPrefix(line, "data: ")
		if data == "[DONE]" {
			break
		}
		var chunk struct {
			Choices []struct {
				Delta struct {
					Content string `json:"content"`
				} `json:"delta"`
			} `json:"choices"`
		}
		if err := json.Unmarshal([]byte(data), &chunk); err != nil || len(chunk.Choices) == 0 {
			continue
		}
		text := chunk.Choices[0].Delta.Content
		if text == "" {
			continue
		}
		full += text
		fmt.Fprintf(w, "data: %s\n\n", sseEscape(text))
		flush(w)
	}
	return full, scanner.Err()
}

// openAIContent is a single content block in an OpenAI message.
type openAIContent struct {
	Type     string          `json:"type"`
	Text     string          `json:"text,omitempty"`
	ImageURL *openAIImageURL `json:"image_url,omitempty"`
}

type openAIImageURL struct {
	URL string `json:"url"`
}

type openAIMessage struct {
	Role    string `json:"role"`
	Content any    `json:"content"` // string OR []openAIContent
}

func (p *OpenAICompatProvider) buildMessages(systemPrompt string, sess *model.Session, history []model.Message) []openAIMessage {
	msgs := []openAIMessage{
		{Role: "system", Content: systemPrompt},
	}

	// Collect image context blocks for the first user message
	var imgBlocks []openAIContent
	if p.info.SupportsImg {
		for _, img := range sess.ResumeImages {
			if img.Data != "" {
				imgBlocks = append(imgBlocks, openAIContent{
					Type:     "image_url",
					ImageURL: &openAIImageURL{URL: "data:" + img.MimeType + ";base64," + img.Data},
				})
			}
		}
		for _, img := range sess.JDImages {
			if img.Data != "" {
				imgBlocks = append(imgBlocks, openAIContent{
					Type:     "image_url",
					ImageURL: &openAIImageURL{URL: "data:" + img.MimeType + ";base64," + img.Data},
				})
			}
		}
	}

	firstUser := true
	for _, m := range history {
		role := "assistant"
		if m.Role == model.RoleUser {
			role = "user"
		}

		// Only merge image blocks into the first user message when provider supports images
		// AND there are actually images. Otherwise always send plain string content.
		if role == "user" && firstUser && p.info.SupportsImg && len(imgBlocks) > 0 {
			blocks := append(imgBlocks, openAIContent{Type: "text", Text: m.Content})
			msgs = append(msgs, openAIMessage{Role: "user", Content: blocks})
			firstUser = false
		} else {
			if role == "user" {
				firstUser = false
			}
			// Always send content as a plain string for non-vision messages
			msgs = append(msgs, openAIMessage{Role: role, Content: m.Content})
		}
	}
	return msgs
}
