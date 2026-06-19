// Package skill 实现有状态的多轮交互技能系统。
//
// Skill 与普通 Tool 的核心区别：
//   - Tool 是无状态的单次调用（如：搜索、计算）
//   - Skill 是有状态的多轮交互（跨多个对话轮次维持上下文）
//
// 架构：SkillContext（共享状态）+ Skill 接口 + SkillRegistry（注册中心）
package skill

// SkillContext 技能执行上下文，在多轮对话间共享。
type SkillContext struct {
	UserID    string
	SessionID string
	Phase     string         // 当前面试阶段
	History   []TurnMessage  // 当前技能的对话历史
	WeakTags  []string       // 用户薄弱点标签
	Metadata  map[string]any // 技能私有状态（如测验的当前题目、已答题数）
}

type TurnMessage struct {
	Role    string // "interviewer" | "candidate"
	Content string
}

// Skill 技能接口——有状态的多轮交互能力单元。
type Skill interface {
	// Name 技能唯一标识。
	Name() string
	// Description 技能说明（用于调试和日志）。
	Description() string
	// Priority 优先级，数字越大越优先匹配（0-100）。
	Priority() int
	// CanActivate 判断当前上下文是否应该激活此技能。
	CanActivate(ctx *SkillContext, trigger string) bool
	// BuildSystemPrompt 返回此技能专属的 system prompt 片段，
	// 追加到主 prompt 后，让 LLM 扮演特定角色。
	BuildSystemPrompt(ctx *SkillContext) string
	// OnTurnEnd 每轮结束后的钩子，用于更新技能内部状态。
	OnTurnEnd(ctx *SkillContext, candidateReply string)
	// IsComplete 技能是否已完成（多轮结束条件）。
	IsComplete(ctx *SkillContext) bool
}
