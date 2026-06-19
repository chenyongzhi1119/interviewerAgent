package agent

import (
	"context"
	"fmt"
	"interviewer-agent/internal/difficulty"
	"interviewer-agent/internal/memory"
	"interviewer-agent/internal/model"
	"interviewer-agent/internal/skill"
)

// BuildSystemPrompt 组装完整 system prompt，集成三个系统的 prompt 注入：
//  1. 公司角色 + 轮次指令（基础）
//  2. 动态难度调度器指令（Phase / 难度档）
//  3. 记忆系统薄弱点指令
//  4. 当前激活 Skill 的专属指令（覆盖普通面试逻辑）
func BuildSystemPrompt(
	profile *model.CompanyProfile,
	session *model.Session,
	memSvc *memory.MemoryService, // 可为 nil（记忆系统未启用时）
	sched *difficulty.PhaseScheduler, // 可为 nil（难度系统未启用时）
	skillReg *skill.SkillRegistry, // 可为 nil
) string {
	round, ok := profile.Rounds[session.Round]
	if !ok {
		round = &model.RoundConfig{
			Title:        fmt.Sprintf("第 %d 轮面试", session.Round),
			Instructions: "请根据候选人的简历和 JD 进行综合面试。",
		}
	}

	base := fmt.Sprintf(`%s

## 当前轮次：%s

### 本轮行为要求
%s

---

## 岗位 JD
%s

---

## 候选人简历
%s

---

## 面试规则
- 每次只提一个问题，等候选人回答后再继续
- 根据候选人回答灵活追问，不要照本宣科
- 始终保持专业、严谨但不刁难的态度
- 全程使用中文`,
		profile.RoleDescription,
		round.Title,
		round.Instructions,
		session.JD,
		session.Resume,
	)

	// ── 注入 1：记忆系统薄弱点 ──────────────────────────────────
	if memSvc != nil && session.UserID != "" {
		base += memSvc.BuildWeaknessPrompt(context.Background(), session.UserID)
	}

	// ── 注入 2：动态难度调度器 ───────────────────────────────────
	if sched != nil && session.DiffPhase != "" {
		st := sessionToDiffState(session)
		var weakTags []string
		if memSvc != nil && session.UserID != "" {
			weakTags = memSvc.GetTopWeakTags(context.Background(), session.UserID, 3)
		}
		base += sched.BuildQuestionPrompt(st, weakTags)
	}

	// ── 注入 3：激活的 Skill 专属 prompt（最高优先级，覆盖普通流程）──
	if skillReg != nil && session.ActiveSkill != "" {
		s := skillReg.Get(session.ActiveSkill)
		if s != nil {
			skillCtx := &skill.SkillContext{
				UserID:    session.UserID,
				SessionID: session.ID,
				Phase:     session.DiffPhase,
				WeakTags:  nil,
				Metadata:  session.SkillMetadata,
			}
			if memSvc != nil && session.UserID != "" {
				skillCtx.WeakTags = memSvc.GetTopWeakTags(context.Background(), session.UserID, 3)
			}
			base += "\n\n" + s.BuildSystemPrompt(skillCtx)
		}
	}

	return base
}

// BuildEvaluationPrompt builds the prompt to trigger round evaluation.
func BuildEvaluationPrompt(profile *model.CompanyProfile, session *model.Session) string {
	round, ok := profile.Rounds[session.Round]
	roundTitle := fmt.Sprintf("第 %d 轮", session.Round)
	if ok {
		roundTitle = round.Title
	}

	diffSummary := ""
	if session.DiffTotalQ > 0 {
		diffSummary = fmt.Sprintf(
			"\n\n【动态难度追踪】本轮共 %d 道题，最终阶段：%s，难度档：%d/5。",
			session.DiffTotalQ, session.DiffPhase, session.DiffLevel,
		)
	}

	return fmt.Sprintf(`本轮面试（%s）已结束。请根据以上完整的对话记录，按照以下评估标准给出评估：%s

%s`, roundTitle, diffSummary, profile.EvaluationRubric)
}

// sessionToDiffState 从 session 恢复 difficulty.State（无需单独存储）。
func sessionToDiffState(session *model.Session) *difficulty.State {
	ph := difficulty.Phase(session.DiffPhase)
	if ph == "" {
		ph = difficulty.PhaseBasic
	}
	d := session.DiffLevel
	if d < 1 {
		d = 2
	}
	return &difficulty.State{
		Phase:            ph,
		Difficulty:       d,
		ConsecCorrect:    session.DiffConsecOK,
		ConsecWrong:      session.DiffConsecFail,
		PhaseQuestionNum: session.DiffPhaseQ,
		TotalQuestions:   session.DiffTotalQ,
		PhaseScores:      make(map[difficulty.Phase][]float64),
	}
}

// diffStateToSession 把调度器状态写回 session。
func diffStateToSession(st *difficulty.State, session *model.Session) {
	session.DiffPhase = string(st.Phase)
	session.DiffLevel = st.Difficulty
	session.DiffConsecOK = st.ConsecCorrect
	session.DiffConsecFail = st.ConsecWrong
	session.DiffPhaseQ = st.PhaseQuestionNum
	session.DiffTotalQ = st.TotalQuestions
}
