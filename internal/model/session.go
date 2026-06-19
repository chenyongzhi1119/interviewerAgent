package model

import "time"

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

type SessionStatus string

const (
	StatusChatting   SessionStatus = "chatting"
	StatusEvaluating SessionStatus = "evaluating"
	StatusDone       SessionStatus = "done"
)

// AttachedFile holds a base64-encoded file (PDF or image) for passing to the LLM.
type AttachedFile struct {
	Data     string `json:"data"`      // base64-encoded file content
	MimeType string `json:"mime_type"` // e.g. "application/pdf", "image/png"
	Name     string `json:"name"`      // display name
}

// Message is a single turn in the interview conversation.
type Message struct {
	Role     Role   `json:"role"`
	Content  string `json:"content"`
	IsHidden bool   `json:"is_hidden,omitempty"` // not shown in UI; used for system openers
}

// Session holds the full state of one interview session.
type Session struct {
	ID           string         `json:"id"`
	Company      string         `json:"company"`
	Round        int            `json:"round"`
	Provider     string         `json:"provider"`
	UserID       string         `json:"user_id,omitempty"` // 关联用户（可选）
	JD           string         `json:"jd"`
	Resume       string         `json:"resume"`
	ResumePDF    *AttachedFile  `json:"resume_pdf,omitempty"`
	ResumeImages []AttachedFile `json:"resume_images,omitempty"`
	JDImages     []AttachedFile `json:"jd_images,omitempty"`
	Messages     []Message      `json:"messages"`
	Status       SessionStatus  `json:"status"`
	CreatedAt    time.Time      `json:"created_at"`

	// 动态难度调度状态（持久化到 session 文件）
	DiffPhase      string  `json:"diff_phase,omitempty"`      // basic/experience/design
	DiffLevel      int     `json:"diff_level,omitempty"`      // 1-5
	DiffConsecOK   int     `json:"diff_consec_ok,omitempty"`  // 连续答对数
	DiffConsecFail int     `json:"diff_consec_fail,omitempty"` // 连续答错数
	DiffPhaseQ     int     `json:"diff_phase_q,omitempty"`    // 当前阶段题数
	DiffTotalQ     int     `json:"diff_total_q,omitempty"`    // 总题数

	// 当前激活的 Skill 名（空=普通面试模式）
	ActiveSkill     string         `json:"active_skill,omitempty"`
	SkillMetadata   map[string]any `json:"skill_metadata,omitempty"`
	SkillTurnCount  int            `json:"skill_turn_count,omitempty"`

	// Runtime-only fields — never serialized to disk
	ProviderKey     string `json:"-"`
	ProviderModel   string `json:"-"`
	ProviderBaseURL string `json:"-"`
}
