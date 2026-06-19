package memory

import "time"

// UserProfile 用户全局画像，跨会话持久化。
type UserProfile struct {
	UserID         string    `json:"user_id"`
	InterviewCount int       `json:"interview_count"`
	OverallLevel   int       `json:"overall_level"` // 1-5，综合评级
	TotalQuestions int       `json:"total_questions"`
	AvgScore       float64   `json:"avg_score"` // 0-100
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// WeaknessRecord 薄弱点记录，低分积累、高分移除、过期淘汰。
type WeaknessRecord struct {
	UserID          string    `json:"user_id"`
	Tag             string    `json:"tag"`           // 技术标签，如 "Redis","分布式锁"
	WeaknessScore   float64   `json:"weakness_score"` // 0-100，越高越薄弱
	OccurrenceCount int       `json:"occurrence_count"`
	LastSeen        time.Time `json:"last_seen"`
	ExpiresAt       time.Time `json:"expires_at"` // 30天未强化则淘汰
}

// QuestionRecord 单次提问记录。
type QuestionRecord struct {
	ID         int64     `json:"id"`
	SessionID  string    `json:"session_id"`
	UserID     string    `json:"user_id"`
	Phase      string    `json:"phase"`      // basic / experience / design
	Difficulty int       `json:"difficulty"` // 1-5
	Tags       []string  `json:"tags"`
	Question   string    `json:"question"`
	Answer     string    `json:"answer"`
	Score      float64   `json:"score"`   // 0-100，由 LLM 评分
	CreatedAt  time.Time `json:"created_at"`
}

// SessionContext 短期记忆：滑动窗口对话上下文。
type SessionContext struct {
	SessionID string
	UserID    string
	Window    []ContextMessage // 最近 N 条对话
	MaxWindow int
}

type ContextMessage struct {
	Role    string // "interviewer" | "candidate"
	Content string
	Phase   string
	Score   float64
}

// WeaknessUpdate 薄弱点更新策略。
type WeaknessUpdate struct {
	Tag   string
	Score float64 // 本次得分
}

const (
	WeaknessThresholdAdd    = 60.0  // 低于此分数 → 记录薄弱点
	WeaknessThresholdRemove = 85.0  // 连续高于此分数 → 移除薄弱点
	WeaknessExpireDays      = 30    // 超过 30 天未出现 → 淘汰
	ContextWindowSize       = 10    // 短期记忆滑动窗口大小
)
