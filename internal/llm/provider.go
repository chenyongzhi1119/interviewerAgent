package llm

import (
	"context"
	"fmt"
	"io"
	"interviewer-agent/internal/model"
)

// Provider is the common interface for all LLM backends.
type Provider interface {
	Stream(ctx context.Context, systemPrompt string, sess *model.Session, history []model.Message, w io.Writer) (string, error)
	// ExtractImageText uses the provider's vision capability to OCR an image.
	// Returns ErrNoVision if the provider/model does not support images.
	ExtractImageText(ctx context.Context, base64Data, mimeType string) (string, error)
	Info() ProviderInfo
}

// ErrNoVision is returned when a provider cannot process images.
var ErrNoVision = fmt.Errorf("此供应商不支持图片识别")

// ProviderInfo describes a provider.
type ProviderInfo struct {
	ID             string `json:"id"`
	DisplayName    string `json:"display_name"`
	Model          string `json:"model"`
	SupportsPDF    bool   `json:"supports_pdf"`
	SupportsImg    bool   `json:"supports_img"`
	IsServerConfig bool   `json:"is_server_config"` // key configured via env var
	RegisterURL    string `json:"register_url"`
	DefaultBaseURL string `json:"default_base_url"`
}

// Message mirrors model.Message (avoids circular import in tests).
type Message = model.Message

// Registry holds providers that were configured at startup via env vars.
type Registry map[string]Provider

// List returns provider infos in display order.
func (r Registry) List() []ProviderInfo {
	order := []string{"anthropic", "openai", "deepseek", "glm", "qwen"}
	var out []ProviderInfo
	for _, id := range order {
		if p, ok := r[id]; ok {
			out = append(out, p.Info())
		}
	}
	for id, p := range r {
		found := false
		for _, o := range order {
			if o == id {
				found = true
				break
			}
		}
		if !found {
			out = append(out, p.Info())
		}
	}
	return out
}

// KnownProviders returns metadata for all supported providers, marking which
// ones are already server-configured.
func KnownProviders(registry Registry) []ProviderInfo {
	all := []ProviderInfo{
		{
			ID: "anthropic", DisplayName: "Claude（Anthropic）",
			Model: "claude-sonnet-4-6", SupportsPDF: true, SupportsImg: true,
			RegisterURL: "https://console.anthropic.com",
		},
		{
			ID: "openai", DisplayName: "GPT-4.1（OpenAI）",
			Model: "gpt-4.1", SupportsPDF: false, SupportsImg: true,
			RegisterURL:    "https://platform.openai.com",
			DefaultBaseURL: "https://api.openai.com/v1",
		},
		{
			ID: "deepseek", DisplayName: "DeepSeek",
			Model: "deepseek-chat", SupportsPDF: false, SupportsImg: false, // deepseek-chat 不支持图片
			RegisterURL:    "https://platform.deepseek.com",
			DefaultBaseURL: "https://api.deepseek.com/v1",
		},
		{
			ID: "glm", DisplayName: "GLM-4（智谱 AI）",
			Model: "glm-4-flash", SupportsPDF: false, SupportsImg: false, // 图片需改用 glm-4v 模型
			RegisterURL:    "https://open.bigmodel.cn",
			DefaultBaseURL: "https://open.bigmodel.cn/api/paas/v4",
		},
		{
			ID: "qwen", DisplayName: "Qwen（阿里云）",
			Model: "qwen-plus", SupportsPDF: false, SupportsImg: true,
			RegisterURL:    "https://dashscope.aliyuncs.com",
			DefaultBaseURL: "https://dashscope.aliyuncs.com/compatible-mode/v1",
		},
	}
	for i, p := range all {
		if _, ok := registry[p.ID]; ok {
			all[i].IsServerConfig = true
		}
	}
	return all
}

// NewProviderFromKey creates a provider on-the-fly using a caller-supplied key.
// id must be one of the known provider presets.
func NewProviderFromKey(id, apiKey, model, baseURL string) Provider {
	if id == "anthropic" {
		return NewAnthropic(apiKey, model)
	}
	return NewOpenAICompat(id, apiKey, baseURL, model)
}
