package difficulty

import "fmt"

// Phase 面试阶段。
type Phase string

const (
	PhaseBasic      Phase = "basic"      // 基础知识
	PhaseExperience Phase = "experience" // 项目经验
	PhaseDesign     Phase = "design"     // 系统设计
)

// State 调度器状态，记录当前阶段、难度和连续答题情况。
type State struct {
	Phase            Phase
	Difficulty       int // 1-5
	ConsecCorrect    int // 连续答对次数
	ConsecWrong      int // 连续答错次数
	PhaseQuestionNum int // 当前阶段已问题数
	TotalQuestions   int // 总问题数
	PhaseScores      map[Phase][]float64 // 各阶段得分记录
}

// PhaseScheduler 三阶段调度器：basic → experience → design，自适应难度。
type PhaseScheduler struct {
	// 各阶段最少问题数，达到后可升阶
	PhaseMinQ map[Phase]int
	// 升档：连续 N 次得分 >= 75 → 难度 +1
	UpThreshold int
	UpScore     float64
	// 降档：连续 N 次得分 < 55 → 难度 -1
	DownThreshold int
	DownScore     float64
}

// DefaultScheduler 默认配置。
func DefaultScheduler() *PhaseScheduler {
	return &PhaseScheduler{
		PhaseMinQ:     map[Phase]int{PhaseBasic: 3, PhaseExperience: 3, PhaseDesign: 2},
		UpThreshold:   2,
		UpScore:       75.0,
		DownThreshold: 2,
		DownScore:     55.0,
	}
}

// NewState 创建初始状态（根据用户历史等级设定起始难度）.
func NewState(overallLevel int) *State {
	d := overallLevel
	if d < 1 {
		d = 1
	}
	if d > 3 {
		d = 3 // 初始难度最高 3，给上升空间
	}
	return &State{
		Phase:       PhaseBasic,
		Difficulty:  d,
		PhaseScores: make(map[Phase][]float64),
	}
}

// RecordScore 记录一次答题得分，返回是否升降了档/阶。
func (s *PhaseScheduler) RecordScore(st *State, score float64) (diffChanged bool, phaseChanged bool) {
	st.TotalQuestions++
	st.PhaseQuestionNum++
	st.PhaseScores[st.Phase] = append(st.PhaseScores[st.Phase], score)

	// 更新连续计数
	if score >= s.UpScore {
		st.ConsecCorrect++
		st.ConsecWrong = 0
	} else if score < s.DownScore {
		st.ConsecWrong++
		st.ConsecCorrect = 0
	} else {
		st.ConsecCorrect = 0
		st.ConsecWrong = 0
	}

	// 难度自适应
	if st.ConsecCorrect >= s.UpThreshold && st.Difficulty < 5 {
		st.Difficulty++
		st.ConsecCorrect = 0
		diffChanged = true
	} else if st.ConsecWrong >= s.DownThreshold && st.Difficulty > 1 {
		st.Difficulty--
		st.ConsecWrong = 0
		diffChanged = true
	}

	// 阶段推进：当前阶段问够了且均分合格 → 升阶
	minQ := s.PhaseMinQ[st.Phase]
	if st.PhaseQuestionNum >= minQ && st.Phase != PhaseDesign {
		avg := avgScores(st.PhaseScores[st.Phase])
		if avg >= 50 { // 均分 50 以上才允许升阶
			st.Phase = nextPhase(st.Phase)
			st.PhaseQuestionNum = 0
			phaseChanged = true
		}
	}
	return
}

// BuildQuestionPrompt 根据当前状态生成给 LLM 的出题指令片段。
func (s *PhaseScheduler) BuildQuestionPrompt(st *State, weakTags []string) string {
	phaseDesc := map[Phase]string{
		PhaseBasic:      "基础知识验证（数据结构、语言特性、网络/操作系统原理）",
		PhaseExperience: "项目经验深挖（设计决策、踩坑经历、性能优化案例）",
		PhaseDesign:     "系统设计与业务落地（从业务需求推导技术架构）",
	}
	diffDesc := map[int]string{
		1: "简单（定义/概念层）",
		2: "基础（原理+简单应用）",
		3: "中等（有一定深度，需要举例）",
		4: "较难（需要权衡取舍或手撕思路）",
		5: "困难（开放性，无标准答案，考察视野）",
	}

	prompt := fmt.Sprintf(
		"\n\n【出题指令】当前阶段：%s | 难度：%s（第 %d 题）。",
		phaseDesc[st.Phase], diffDesc[st.Difficulty], st.PhaseQuestionNum+1,
	)
	if len(weakTags) > 0 {
		prompt += fmt.Sprintf("请优先围绕候选人薄弱点【%s】出题。", joinTags(weakTags))
	}
	prompt += "每次只问一个问题，等候选人回答后再继续。"
	return prompt
}

// Summary 生成阶段总结（用于评估 prompt）。
func (s *PhaseScheduler) Summary(st *State) string {
	var parts []string
	for _, ph := range []Phase{PhaseBasic, PhaseExperience, PhaseDesign} {
		scores := st.PhaseScores[ph]
		if len(scores) == 0 {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s 阶段均分 %.0f（%d 题）", ph, avgScores(scores), len(scores)))
	}
	result := fmt.Sprintf("本次面试共 %d 题，峰值难度 %d 档", st.TotalQuestions, st.Difficulty)
	if len(parts) > 0 {
		result += "：" + joinTags(parts)
	}
	return result
}

func nextPhase(p Phase) Phase {
	switch p {
	case PhaseBasic:
		return PhaseExperience
	case PhaseExperience:
		return PhaseDesign
	default:
		return PhaseDesign
	}
}

func avgScores(scores []float64) float64 {
	if len(scores) == 0 {
		return 0
	}
	var sum float64
	for _, s := range scores {
		sum += s
	}
	return sum / float64(len(scores))
}

func joinTags(tags []string) string {
	result := ""
	for i, t := range tags {
		if i > 0 {
			result += "、"
		}
		result += t
	}
	return result
}
