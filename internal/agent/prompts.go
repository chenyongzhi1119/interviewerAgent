package agent

import (
	"fmt"
	"interviewer-agent/internal/model"
)

// BuildSystemPrompt assembles the full system prompt for the current interview state.
func BuildSystemPrompt(profile *model.CompanyProfile, session *model.Session) string {
	round, ok := profile.Rounds[session.Round]
	if !ok {
		round = &model.RoundConfig{
			Title:        fmt.Sprintf("第 %d 轮面试", session.Round),
			Instructions: "请根据候选人的简历和 JD 进行综合面试。",
		}
	}

	return fmt.Sprintf(`%s

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
}

// BuildEvaluationPrompt builds the prompt to trigger round evaluation.
func BuildEvaluationPrompt(profile *model.CompanyProfile, session *model.Session) string {
	round, ok := profile.Rounds[session.Round]
	roundTitle := fmt.Sprintf("第 %d 轮", session.Round)
	if ok {
		roundTitle = round.Title
	}

	return fmt.Sprintf(`本轮面试（%s）已结束。请根据以上完整的对话记录，按照以下评估标准给出评估：

%s`, roundTitle, profile.EvaluationRubric)
}
