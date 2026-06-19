package skill

import (
	"sort"
	"sync"
)

// SkillRegistry 技能注册中心，支持按优先级匹配、插拔扩展。
//
// 使用方式：
//
//	registry := skill.NewRegistry()
//	registry.Register(&QuizSkill{})
//	registry.Register(&TeachSkill{})
//	matched := registry.Match(ctx, userInput)
type SkillRegistry struct {
	mu     sync.RWMutex
	skills []Skill // 按 Priority 降序排列
}

func NewRegistry() *SkillRegistry {
	r := &SkillRegistry{}
	// 注册 4 个内置 Skill
	r.Register(&QuizSkill{})
	r.Register(&TeachSkill{})
	r.Register(&ProjectHighlightSkill{})
	r.Register(&CompareSkill{})
	return r
}

// Register 注册一个 Skill，自动按优先级排序。
func (r *SkillRegistry) Register(s Skill) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.skills = append(r.skills, s)
	sort.Slice(r.skills, func(i, j int) bool {
		return r.skills[i].Priority() > r.skills[j].Priority()
	})
}

// Match 按优先级顺序遍历，返回第一个 CanActivate 为真的 Skill。
// 若无匹配则返回 nil（走普通面试逻辑）。
func (r *SkillRegistry) Match(ctx *SkillContext, trigger string) Skill {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, s := range r.skills {
		if s.CanActivate(ctx, trigger) {
			return s
		}
	}
	return nil
}

// Get 按名称获取已注册的 Skill。
func (r *SkillRegistry) Get(name string) Skill {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, s := range r.skills {
		if s.Name() == name {
			return s
		}
	}
	return nil
}

// List 返回所有已注册 Skill 的元信息。
func (r *SkillRegistry) List() []SkillInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var infos []SkillInfo
	for _, s := range r.skills {
		infos = append(infos, SkillInfo{
			Name:        s.Name(),
			Description: s.Description(),
			Priority:    s.Priority(),
		})
	}
	return infos
}

type SkillInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}
