package extract

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const ocrPrompt = "请提取图片中的所有文字内容，原样输出，保留原有段落格式，不要添加任何解释或额外内容。"

// VisionConfig specifies which vision provider to use for image OCR.
type VisionConfig struct {
	ProviderID  string
	APIKey      string
	Model       string // empty = use default
	BaseURL     string // empty = use default
}

var visionDefaults = map[string]struct {
	baseURL string
	model   string
}{
	"anthropic": {"https://api.anthropic.com", "claude-sonnet-4-6"},
	"openai":    {"https://api.openai.com/v1", "gpt-4o"},
	"qwen":      {"https://dashscope.aliyuncs.com/compatible-mode/v1", "qwen-vl-plus"},
}

// ImageOCR calls a vision LLM to extract text from a base64-encoded image.
func ImageOCR(ctx context.Context, base64Data, mimeType string, cfg VisionConfig) (string, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if cfg.ProviderID == "anthropic" {
		return ocrViaAnthropic(ctx, base64Data, mimeType, cfg)
	}
	return ocrViaOpenAICompat(ctx, base64Data, mimeType, cfg)
}

// ─── Anthropic ───────────────────────────────────────────────────────────────

func ocrViaAnthropic(ctx context.Context, base64Data, mimeType string, cfg VisionConfig) (string, error) {
	baseURL := cfg.BaseURL
	model := cfg.Model
	if d, ok := visionDefaults["anthropic"]; ok {
		if baseURL == "" {
			baseURL = d.baseURL
		}
		if model == "" {
			model = d.model
		}
	}

	body := map[string]any{
		"model":      model,
		"max_tokens": 4096,
		"messages": []map[string]any{
			{
				"role": "user",
				"content": []map[string]any{
					{
						"type": "image",
						"source": map[string]any{
							"type":       "base64",
							"media_type": mimeType,
							"data":       base64Data,
						},
					},
					{"type": "text", "text": ocrPrompt},
				},
			},
		},
	}

	resp, err := doPost(ctx, baseURL+"/v1/messages", cfg.APIKey, body, map[string]string{
		"anthropic-version": "2023-06-01",
		"x-api-key":         cfg.APIKey,
		"authorization":     "", // Anthropic uses x-api-key, not Bearer
	})
	if err != nil {
		return "", err
	}

	var result struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return "", fmt.Errorf("parse response: %w", err)
	}
	if result.Error != nil {
		return "", fmt.Errorf("API error: %s", result.Error.Message)
	}
	for _, c := range result.Content {
		if c.Type == "text" {
			return strings.TrimSpace(c.Text), nil
		}
	}
	return "", fmt.Errorf("no text in response")
}

// ─── OpenAI-compatible (GPT, Qwen, etc.) ─────────────────────────────────────

func ocrViaOpenAICompat(ctx context.Context, base64Data, mimeType string, cfg VisionConfig) (string, error) {
	d, ok := visionDefaults[cfg.ProviderID]
	baseURL := cfg.BaseURL
	model := cfg.Model
	if ok {
		if baseURL == "" {
			baseURL = d.baseURL
		}
		if model == "" {
			model = d.model
		}
	}
	if baseURL == "" {
		return "", fmt.Errorf("base URL not configured for provider %q", cfg.ProviderID)
	}

	body := map[string]any{
		"model": model,
		"messages": []map[string]any{
			{
				"role": "user",
				"content": []map[string]any{
					{
						"type":      "image_url",
						"image_url": map[string]string{"url": "data:" + mimeType + ";base64," + base64Data},
					},
					{"type": "text", "text": ocrPrompt},
				},
			},
		},
		"stream": false,
	}

	resp, err := doPost(ctx, strings.TrimRight(baseURL, "/")+"/chat/completions", cfg.APIKey, body, nil)
	if err != nil {
		return "", err
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return "", fmt.Errorf("parse response: %w", err)
	}
	if result.Error != nil {
		return "", fmt.Errorf("API error: %s", result.Error.Message)
	}
	if len(result.Choices) == 0 {
		return "", fmt.Errorf("empty response")
	}
	return strings.TrimSpace(result.Choices[0].Message.Content), nil
}

// ─── shared HTTP helper ───────────────────────────────────────────────────────

func doPost(ctx context.Context, url, apiKey string, body any, extraHeaders map[string]string) ([]byte, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	// Default: Bearer auth (OpenAI-compat)
	if _, hasAuth := extraHeaders["authorization"]; !hasAuth {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}
	for k, v := range extraHeaders {
		if v != "" {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s %s", resp.Status, string(data))
	}
	return data, nil
}
