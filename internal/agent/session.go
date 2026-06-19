package agent

import (
	"context"
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
	"interviewer-agent/internal/difficulty"
	"interviewer-agent/internal/llm"
	"interviewer-agent/internal/memory"
	"interviewer-agent/internal/model"
	"interviewer-agent/internal/skill"
)

// Store manages active interview sessions, company profiles, and LLM providers.
type Store struct {
	mu          sync.RWMutex
	sessions    map[string]*model.Session
	companies   map[string]*model.CompanyProfile
	providers   llm.Registry
	sessionsDir string

	// 三大增强系统（可选，nil = 不启用）
	MemSvc   *memory.MemoryService
	Sched    *difficulty.PhaseScheduler
	SkillReg *skill.SkillRegistry
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

	systemPrompt := BuildSystemPrompt(profile, sess, s.MemSvc, s.Sched, s.SkillReg)
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

	// ── Skill 状态机：检测触发 / 推进当前 Skill ──────────────────
	if s.SkillReg != nil {
		skillCtx := &skill.SkillContext{
			UserID:    sess.UserID,
			SessionID: sess.ID,
			Phase:     sess.DiffPhase,
			Metadata:  sess.SkillMetadata,
		}
		if sess.SkillMetadata == nil {
			sess.SkillMetadata = make(map[string]any)
			skillCtx.Metadata = sess.SkillMetadata
		}

		// 如果有激活的 Skill，检查是否完成
		if sess.ActiveSkill != "" {
			sk := s.SkillReg.Get(sess.ActiveSkill)
			if sk != nil {
				if sk.IsComplete(skillCtx) {
					// Skill 完成，退出技能模式
					sess.ActiveSkill = ""
					sess.SkillMetadata = nil
				} else {
					// 继续技能模式
					sk.OnTurnEnd(skillCtx, userMsg)
					sess.SkillTurnCount++
				}
			}
		} else {
			// 未激活 Skill，检测是否触发新 Skill
			if matched := s.SkillReg.Match(skillCtx, userMsg); matched != nil {
				sess.ActiveSkill = matched.Name()
				sess.SkillMetadata = map[string]any{}
				sess.SkillTurnCount = 0
			}
		}
	}

	sess.Messages = append(sess.Messages, model.Message{Role: model.RoleUser, Content: userMsg})
	systemPrompt := BuildSystemPrompt(profile, sess, s.MemSvc, s.Sched, s.SkillReg)
	reply, err := provider.Stream(nil, systemPrompt, sess, sess.Messages, w)
	if err != nil {
		return err
	}
	sess.Messages = append(sess.Messages, model.Message{Role: model.RoleAssistant, Content: reply})

	// ── 难度自适应：LLM 内置评分提示（简化版，实际可由 LLM 显式输出分数）──
	// 这里用回复长度作为代理信号（详细回答 → 得分较高）
	// 生产环境可在 LLM system prompt 中要求输出 JSON score 字段
	if s.Sched != nil && sess.DiffPhase != "" {
		score := estimateScore(reply)
		st := sessionToDiffState(sess)
		diffChanged, phaseChanged := s.Sched.RecordScore(st, score)
		diffStateToSession(st, sess)

		if s.MemSvc != nil && sess.UserID != "" {
			s.MemSvc.RecordAnswer(context.Background(), &memory.QuestionRecord{
				SessionID:  sess.ID,
				UserID:     sess.UserID,
				Phase:      string(st.Phase),
				Difficulty: sess.DiffLevel,
				Question:   userMsg,
				Score:      score,
			})
		}

		if diffChanged {
			log.Printf("session %s: difficulty → %d", sess.ID, sess.DiffLevel)
		}
		if phaseChanged {
			log.Printf("session %s: phase → %s", sess.ID, sess.DiffPhase)
		}
	}

	s.persist(sess)
	return nil
}

// estimateScore 用回复质量代理估算候选人得分（0-100）。
// 生产环境可要求 LLM 在 system prompt 中显式输出 "score: XX" 然后解析。
func estimateScore(candidateReply string) float64 {
	n := len([]rune(candidateReply))
	switch {
	case n < 20:
		return 30 // 太短，基本没答
	case n < 80:
		return 55
	case n < 200:
		return 68
	case n < 500:
		return 75
	default:
		return 82 // 详细作答，基础分较高
	}
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
	systemPrompt := BuildSystemPrompt(profile, sess, s.MemSvc, s.Sched, s.SkillReg)

	reply, err := provider.Stream(nil, systemPrompt, sess, allMsgs, w)
	if err != nil {
		return err
	}

	sess.Messages = append(sess.Messages, model.Message{Role: model.RoleAssistant, Content: reply})
	sess.Status = model.StatusDone
	s.persist(sess)
	return nil
}
