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
	JD           string         `json:"jd"`
	Resume       string         `json:"resume"`
	ResumePDF    *AttachedFile  `json:"resume_pdf,omitempty"`
	ResumeImages []AttachedFile `json:"resume_images,omitempty"`
	JDImages     []AttachedFile `json:"jd_images,omitempty"`
	Messages     []Message      `json:"messages"`
	Status       SessionStatus  `json:"status"`
	CreatedAt    time.Time      `json:"created_at"`

	// Runtime-only fields — never serialized to disk
	ProviderKey     string `json:"-"`
	ProviderModel   string `json:"-"`
	ProviderBaseURL string `json:"-"`
}
