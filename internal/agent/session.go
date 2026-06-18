package agent

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
	"interviewer-agent/internal/llm"
	"interviewer-agent/internal/model"
)

// Store manages active interview sessions, company profiles, and LLM providers.
type Store struct {
	mu          sync.RWMutex
	sessions    map[string]*model.Session
	companies   map[string]*model.CompanyProfile
	providers   llm.Registry
	sessionsDir string
}

// NewStore initializes the store.
func NewStore(providers llm.Registry, companiesDir, sessionsDir string) (*Store, error) {
	s := &Store{
		sessions:    make(map[string]*model.Session),
		companies:   make(map[string]*model.CompanyProfile),
		providers:   providers,
		sessionsDir: sessionsDir,
	}
	if err := s.loadCompanies(companiesDir); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(sessionsDir, 0o755); err != nil {
		return nil, fmt.Errorf("create sessions dir: %w", err)
	}
	s.loadPersistedSessions()
	return s, nil
}

func (s *Store) loadCompanies(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read companies dir: %w", err)
	}
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".yaml" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			return err
		}
		var p model.CompanyProfile
		if err := yaml.Unmarshal(data, &p); err != nil {
			return fmt.Errorf("parse %s: %w", e.Name(), err)
		}
		s.companies[p.Name] = &p
	}
	return nil
}

func (s *Store) loadPersistedSessions() {
	entries, err := os.ReadDir(s.sessionsDir)
	if err != nil {
		return
	}
	n := 0
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".json" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.sessionsDir, e.Name()))
		if err != nil {
			continue
		}
		var sess model.Session
		if err := json.Unmarshal(data, &sess); err != nil {
			log.Printf("warn: invalid session file %s: %v", e.Name(), err)
			continue
		}
		s.sessions[sess.ID] = &sess
		n++
	}
	if n > 0 {
		log.Printf("Loaded %d persisted session(s)", n)
	}
}

func (s *Store) persist(sess *model.Session) {
	data, err := json.MarshalIndent(sess, "", "  ")
	if err != nil {
		return
	}
	path := filepath.Join(s.sessionsDir, sess.ID+".json")
	if err := os.WriteFile(path, data, 0o644); err != nil {
		log.Printf("warn: write session %s: %v", sess.ID, err)
	}
}

func (s *Store) Companies() map[string]*model.CompanyProfile {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.companies
}

func (s *Store) Providers() llm.Registry { return s.providers }

func (s *Store) ListSessions() []*model.Session {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Session, 0, len(s.sessions))
	for _, sess := range s.sessions {
		list = append(list, sess)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreatedAt.After(list[j].CreatedAt)
	})
	return list
}

func (s *Store) CreateSession(company, providerID string, round int, jd, resume string, resumePDF *model.AttachedFile, resumeImages, jdImages []model.AttachedFile, providerKey, providerModel, providerBaseURL string) (*model.Session, error) {
	s.mu.RLock()
	_, hasCompany := s.companies[company]
	_, hasProvider := s.providers[providerID]
	s.mu.RUnlock()

	if !hasCompany {
		return nil, fmt.Errorf("unknown company: %s", company)
	}
	// Provider must be either server-configured OR caller supplies a key
	if !hasProvider && providerKey == "" {
		return nil, fmt.Errorf("provider %q is not configured — please enter an API key in settings", providerID)
	}

	sess := &model.Session{
		ID:              uuid.NewString(),
		Company:         company,
		Provider:        providerID,
		Round:           round,
		JD:              jd,
		Resume:          resume,
		ResumePDF:       resumePDF,
		ResumeImages:    resumeImages,
		JDImages:        jdImages,
		Messages:        []model.Message{},
		Status:          model.StatusChatting,
		CreatedAt:       time.Now(),
		ProviderKey:     providerKey,
		ProviderModel:   providerModel,
		ProviderBaseURL: providerBaseURL,
	}
	s.mu.Lock()
	s.sessions[sess.ID] = sess
	s.mu.Unlock()
	s.persist(sess)
	return sess, nil
}

func (s *Store) GetSession(id string) (*model.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sess, ok := s.sessions[id]
	if !ok {
		return nil, fmt.Errorf("session not found: %s", id)
	}
	return sess, nil
}

func (s *Store) getProvider(sess *model.Session) (llm.Provider, error) {
	// Prefer caller-supplied key (entered via UI)
	if sess.ProviderKey != "" {
		return llm.NewProviderFromKey(sess.Provider, sess.ProviderKey, sess.ProviderModel, sess.ProviderBaseURL), nil
	}
	// Fall back to server-configured provider
	p, ok := s.providers[sess.Provider]
	if !ok {
		return nil, fmt.Errorf("provider %q not available — please enter an API key in settings", sess.Provider)
	}
	return p, nil
}

func (s *Store) StartInterview(sess *model.Session, w io.Writer) error {
	s.mu.RLock()
	profile := s.companies[sess.Company]
	s.mu.RUnlock()

	provider, err := s.getProvider(sess)
	if err != nil {
		return err
	}

	round := profile.Rounds[sess.Round]
	roundTitle := fmt.Sprintf("第 %d 轮", sess.Round)
	if round != nil {
		roundTitle = round.Title
	}

	openerText := fmt.Sprintf("请开始「%s」面试，根据候选人的简历和 JD 提出第一个面试问题。直接提问，不要说开场白。", roundTitle)
	opener := model.Message{Role: model.RoleUser, Content: openerText, IsHidden: true}
	sess.Messages = append(sess.Messages, opener)

	systemPrompt := BuildSystemPrompt(profile, sess)
	reply, err := provider.Stream(nil, systemPrompt, sess, sess.Messages, w)
	if err != nil {
		return err
	}

	sess.Messages = append(sess.Messages, model.Message{Role: model.RoleAssistant, Content: reply})
	s.persist(sess)
	return nil
}

func (s *Store) Chat(sess *model.Session, userMsg string, w io.Writer) error {
	s.mu.RLock()
	profile := s.companies[sess.Company]
	s.mu.RUnlock()

	provider, err := s.getProvider(sess)
	if err != nil {
		return err
	}

	sess.Messages = append(sess.Messages, model.Message{Role: model.RoleUser, Content: userMsg})
	systemPrompt := BuildSystemPrompt(profile, sess)
	reply, err := provider.Stream(nil, systemPrompt, sess, sess.Messages, w)
	if err != nil {
		return err
	}

	sess.Messages = append(sess.Messages, model.Message{Role: model.RoleAssistant, Content: reply})
	s.persist(sess)
	return nil
}

func (s *Store) Evaluate(sess *model.Session, w io.Writer) error {
	s.mu.RLock()
	profile := s.companies[sess.Company]
	s.mu.RUnlock()

	provider, err := s.getProvider(sess)
	if err != nil {
		return err
	}

	sess.Status = model.StatusEvaluating
	evalPrompt := BuildEvaluationPrompt(profile, sess)
	allMsgs := append(sess.Messages, model.Message{Role: model.RoleUser, Content: evalPrompt})
	systemPrompt := BuildSystemPrompt(profile, sess)

	reply, err := provider.Stream(nil, systemPrompt, sess, allMsgs, w)
	if err != nil {
		return err
	}

	sess.Messages = append(sess.Messages, model.Message{Role: model.RoleAssistant, Content: reply})
	sess.Status = model.StatusDone
	s.persist(sess)
	return nil
}
