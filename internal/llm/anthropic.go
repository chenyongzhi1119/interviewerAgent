package llm

import (
	"context"
	"fmt"
	"io"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"interviewer-agent/internal/model"
)

// AnthropicProvider calls the Anthropic Messages API.
// It supports PDF documents and images natively.
type AnthropicProvider struct {
	api   anthropic.Client
	model anthropic.Model
}

func NewAnthropic(apiKey, modelID string) *AnthropicProvider {
	if modelID == "" {
		modelID = "claude-sonnet-4-6"
	}
	return &AnthropicProvider{
		api:   anthropic.NewClient(option.WithAPIKey(apiKey)),
		model: anthropic.Model(modelID),
	}
}

func (p *AnthropicProvider) ExtractImageText(ctx context.Context, base64Data, mimeType string) (string, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	stream := p.api.Messages.NewStreaming(ctx, anthropic.MessageNewParams{
		Model:     p.model,
		MaxTokens: 4096,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(
				anthropic.NewImageBlockBase64(mimeType, base64Data),
				anthropic.NewTextBlock("请提取图片中的所有文字内容，原样输出，保留原有段落格式，不要添加任何解释或额外内容。"),
			),
		},
	})
	var full string
	for stream.Next() {
		event := stream.Current()
		if e, ok := event.AsAny().(anthropic.ContentBlockDeltaEvent); ok {
			if delta, ok := e.Delta.AsAny().(anthropic.TextDelta); ok {
				full += delta.Text
			}
		}
	}
	if err := stream.Err(); err != nil {
		return "", err
	}
	return full, nil
}

func (p *AnthropicProvider) Info() ProviderInfo {
	return ProviderInfo{
		ID:          "anthropic",
		DisplayName: "Claude（Anthropic）",
		Model:       string(p.model),
		SupportsPDF: true,
		SupportsImg: true,
	}
}

func (p *AnthropicProvider) Stream(ctx context.Context, systemPrompt string, sess *model.Session, history []model.Message, w io.Writer) (string, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	msgs := buildAnthropicMessages(sess, history)

	stream := p.api.Messages.NewStreaming(ctx, anthropic.MessageNewParams{
		Model:     p.model,
		MaxTokens: 2048,
		System:    []anthropic.TextBlockParam{{Text: systemPrompt}},
		Messages:  msgs,
	})

	var full string
	for stream.Next() {
		event := stream.Current()
		switch e := event.AsAny().(type) {
		case anthropic.ContentBlockDeltaEvent:
			if delta, ok := e.Delta.AsAny().(anthropic.TextDelta); ok {
				chunk := delta.Text
				full += chunk
				fmt.Fprintf(w, "data: %s\n\n", sseEscape(chunk))
				flush(w)
			}
		}
	}
	if err := stream.Err(); err != nil {
		return full, fmt.Errorf("stream error: %w", err)
	}
	return full, nil
}

// buildAnthropicMessages converts session history to Anthropic message params,
// injecting PDF/image context blocks into the first user message.
func buildAnthropicMessages(sess *model.Session, history []model.Message) []anthropic.MessageParam {
	var contextBlocks []anthropic.ContentBlockParamUnion

	hasResumePDF := sess.ResumePDF != nil && sess.ResumePDF.Data != ""
	hasResumeImgs := len(sess.ResumeImages) > 0
	if hasResumePDF || hasResumeImgs {
		contextBlocks = append(contextBlocks, anthropic.NewTextBlock("【候选人简历】"))
		if hasResumePDF {
			contextBlocks = append(contextBlocks,
				anthropic.NewDocumentBlock(anthropic.Base64PDFSourceParam{Data: sess.ResumePDF.Data}),
			)
		}
		for _, img := range sess.ResumeImages {
			if img.Data != "" {
				contextBlocks = append(contextBlocks, anthropic.NewImageBlockBase64(img.MimeType, img.Data))
			}
		}
	}
	if len(sess.JDImages) > 0 {
		contextBlocks = append(contextBlocks, anthropic.NewTextBlock("【岗位 JD 截图】"))
		for _, img := range sess.JDImages {
			if img.Data != "" {
				contextBlocks = append(contextBlocks, anthropic.NewImageBlockBase64(img.MimeType, img.Data))
			}
		}
	}

	var msgs []anthropic.MessageParam
	firstUser := true
	for _, m := range history {
		if m.Role == model.RoleUser {
			var blocks []anthropic.ContentBlockParamUnion
			if firstUser && len(contextBlocks) > 0 {
				blocks = append(blocks, contextBlocks...)
			}
			firstUser = false
			blocks = append(blocks, anthropic.NewTextBlock(m.Content))
			msgs = append(msgs, anthropic.NewUserMessage(blocks...))
		} else {
			msgs = append(msgs, anthropic.NewAssistantMessage(anthropic.NewTextBlock(m.Content)))
		}
	}
	return msgs
}
