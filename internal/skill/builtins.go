package skill

import "strings"

// ── Skill 1: QuizSkill 快速测验 ───────────────────────────────────────────────
//
// 触发词：用户说「来个测试」「快速考一下」「quiz」等
// 状态：当前题号、总题数
// 行为：连续出 5 道选择/判断题，每题立即反馈，最后给总分

type QuizSkill struct{}

func (s *QuizSkill) Name() string        { return "quick_quiz" }
func (s *QuizSkill) Description() string { return "快速测验：连出 5 道题，即时反馈，最后评分" }
func (s *QuizSkill) Priority() int       { return 80 }

func (s *QuizSkill) CanActivate(ctx *SkillContext, trigger string) bool {
	keywords := []string{"测验", "测试一下", "quiz", "来几道题", "快速考", "刷题"}
	lower := strings.ToLower(trigger)
	for _, kw := range keywords {
		if strings.Contains(lower, kw) {
			return true
		}
	}
	return false
}

func (s *QuizSkill) BuildSystemPrompt(ctx *SkillContext) string {
	q := s.currentQ(ctx)
	total := 5
	weakStr := ""
	if len(ctx.WeakTags) > 0 {
		weakStr = "，重点围绕候选人薄弱点：" + strings.Join(ctx.WeakTags, "、")
	}
	return `【快速测验模式】你正在主持一场快速测验` + weakStr + `。
规则：
1. 连续出` + itoa(total) + `道题（当前第` + itoa(q) + `题）
2. 每道题给出 4 个选项（A/B/C/D），等候选人回答
3. 候选人回答后立即告知对错及简要解释
4. 第` + itoa(total) + `题结束后给出总分（X/` + itoa(total) + `）和简短总结
当前面试阶段：` + ctx.Phase + `，请出第` + itoa(q) + `题。`
}

func (s *QuizSkill) OnTurnEnd(ctx *SkillContext, _ string) {
	ctx.Metadata["quiz_q"] = s.currentQ(ctx) + 1
}

func (s *QuizSkill) IsComplete(ctx *SkillContext) bool {
	return s.currentQ(ctx) > 5
}

func (s *QuizSkill) currentQ(ctx *SkillContext) int {
	if ctx.Metadata == nil {
		return 1
	}
	if q, ok := ctx.Metadata["quiz_q"].(int); ok {
		return q
	}
	return 1
}

// ── Skill 2: TeachSkill 概念教学 ─────────────────────────────────────────────
//
// 触发词：「解释一下」「我不太懂」「教我」「什么是 X」
// 状态：教学轮次、当前概念
// 行为：Socratic 式教学，先问候选人理解，再针对性讲解，最后确认掌握

type TeachSkill struct{}

func (s *TeachSkill) Name() string        { return "concept_teach" }
func (s *TeachSkill) Description() string { return "概念教学：Socratic 问答，引导候选人从已知推出未知" }
func (s *TeachSkill) Priority() int       { return 70 }

func (s *TeachSkill) CanActivate(ctx *SkillContext, trigger string) bool {
	keywords := []string{"解释", "不太懂", "教我", "什么是", "帮我理解", "could you explain", "不明白"}
	lower := strings.ToLower(trigger)
	for _, kw := range keywords {
		if strings.Contains(lower, kw) {
			return true
		}
	}
	return false
}

func (s *TeachSkill) BuildSystemPrompt(ctx *SkillContext) string {
	round := s.round(ctx)
	return `【概念教学模式】候选人对某个概念有疑问，请用 Socratic 教学法引导。
步骤（当前第` + itoa(round) + `轮）：
1. 先问候选人目前对这个概念知道哪些（了解已有认知基础）
2. 基于候选人的回答，针对性讲解，类比生活场景
3. 出一道小练习题验证是否理解
4. 候选人回答正确则结束教学，返回正式面试
请用鼓励性语气，避免直接给答案，多用「你觉得呢？」「如果换个场景……」。`
}

func (s *TeachSkill) OnTurnEnd(ctx *SkillContext, reply string) {
	ctx.Metadata["teach_round"] = s.round(ctx) + 1
	if strings.Contains(reply, "理解了") || strings.Contains(reply, "明白了") {
		ctx.Metadata["teach_done"] = true
	}
}

func (s *TeachSkill) IsComplete(ctx *SkillContext) bool {
	done, _ := ctx.Metadata["teach_done"].(bool)
	return done || s.round(ctx) > 4
}

func (s *TeachSkill) round(ctx *SkillContext) int {
	if ctx.Metadata == nil {
		return 1
	}
	if r, ok := ctx.Metadata["teach_round"].(int); ok {
		return r
	}
	return 1
}

// ── Skill 3: ProjectHighlightSkill 项目亮点提炼 ──────────────────────────────
//
// 触发词：「帮我提炼项目亮点」「怎么介绍这个项目」「STAR 法则」
// 状态：提炼进度（收集信息 → 提炼 → 优化）
// 行为：引导候选人描述项目，提炼 STAR 结构亮点，给出面试话术建议

type ProjectHighlightSkill struct{}

func (s *ProjectHighlightSkill) Name() string        { return "project_highlight" }
func (s *ProjectHighlightSkill) Description() string { return "项目亮点提炼：引导候选人用 STAR 法则重构项目描述" }
func (s *ProjectHighlightSkill) Priority() int       { return 60 }

func (s *ProjectHighlightSkill) CanActivate(ctx *SkillContext, trigger string) bool {
	keywords := []string{"项目亮点", "怎么介绍", "star法则", "star原则", "提炼", "怎么说这个项目"}
	lower := strings.ToLower(trigger)
	for _, kw := range keywords {
		if strings.Contains(lower, kw) {
			return true
		}
	}
	return false
}

func (s *ProjectHighlightSkill) BuildSystemPrompt(ctx *SkillContext) string {
	step := s.step(ctx)
	steps := map[int]string{
		1: `请先引导候选人描述项目：问「这个项目的背景是什么？解决了什么业务问题？」`,
		2: `候选人已描述背景，现在引导他说技术方案：问「你具体做了什么？遇到了什么技术难点？」`,
		3: `现在引导他说结果：问「这个项目上线后效果如何？有没有量化的指标？」`,
		4: `信息收集完毕，现在帮他提炼 STAR 结构（Situation/Task/Action/Result），给出 2-3 个可以在面试中突出的亮点，并给出话术建议（30秒版本和2分钟版本）。`,
	}
	desc := steps[step]
	if desc == "" {
		desc = steps[4]
	}
	return `【项目亮点提炼模式】` + desc + `
注意：用引导性问题，不要替候选人总结，让他自己表达，你来优化结构。`
}

func (s *ProjectHighlightSkill) OnTurnEnd(ctx *SkillContext, _ string) {
	ctx.Metadata["ph_step"] = s.step(ctx) + 1
}

func (s *ProjectHighlightSkill) IsComplete(ctx *SkillContext) bool {
	return s.step(ctx) > 4
}

func (s *ProjectHighlightSkill) step(ctx *SkillContext) int {
	if ctx.Metadata == nil {
		return 1
	}
	if st, ok := ctx.Metadata["ph_step"].(int); ok {
		return st
	}
	return 1
}

// ── Skill 4: CompareSkill 技术对比 ───────────────────────────────────────────
//
// 触发词：「X 和 Y 区别」「对比一下」「哪个更好」「vs」
// 状态：对比维度进度
// 行为：引导候选人从多维度对比两个技术，最后给出选型建议框架

type CompareSkill struct{}

func (s *CompareSkill) Name() string        { return "tech_compare" }
func (s *CompareSkill) Description() string { return "技术对比：多维度对比两个技术方案，引导出选型思维框架" }
func (s *CompareSkill) Priority() int       { return 50 }

func (s *CompareSkill) CanActivate(ctx *SkillContext, trigger string) bool {
	keywords := []string{"区别", "对比", "哪个好", " vs ", "和…的区别", "和...区别", "两者"}
	lower := strings.ToLower(trigger)
	for _, kw := range keywords {
		if strings.Contains(lower, kw) {
			return true
		}
	}
	return false
}

func (s *CompareSkill) BuildSystemPrompt(ctx *SkillContext) string {
	dim := s.dim(ctx)
	dimensions := []string{"使用场景/定位", "性能特性（时间/空间/吞吐量）", "一致性/可靠性保证", "运维复杂度", "实际选型建议（给出3个判断标准）"}
	current := "使用场景"
	if dim < len(dimensions) {
		current = dimensions[dim]
	}
	return `【技术对比模式】你在帮候选人深度理解两个技术方案的对比。
当前对比维度（第` + itoa(dim+1) + `/` + itoa(len(dimensions)) + `维）：【` + current + `】
步骤：
1. 先问候选人对「` + current + `」这个维度怎么看
2. 候选人回答后，补充他遗漏的关键点，纠正错误认知
3. 给出一个实际场景例子加深理解
4. 全部维度结束后，帮候选人总结「面试时如何回答对比题的通用框架」。`
}

func (s *CompareSkill) OnTurnEnd(ctx *SkillContext, _ string) {
	ctx.Metadata["cmp_dim"] = s.dim(ctx) + 1
}

func (s *CompareSkill) IsComplete(ctx *SkillContext) bool {
	return s.dim(ctx) >= 5
}

func (s *CompareSkill) dim(ctx *SkillContext) int {
	if ctx.Metadata == nil {
		return 0
	}
	if d, ok := ctx.Metadata["cmp_dim"].(int); ok {
		return d
	}
	return 0
}

// ── helpers ──────────────────────────────────────────────────────────────────

func itoa(n int) string {
	if n < 0 {
		return "0"
	}
	digits := []byte("0123456789")
	if n == 0 {
		return "0"
	}
	var buf [20]byte
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = digits[n%10]
		n /= 10
	}
	return string(buf[pos:])
}
