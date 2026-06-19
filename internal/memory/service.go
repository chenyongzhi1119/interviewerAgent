package memory

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
)

// MemoryService 是记忆系统的统一入口，Cache-Aside 模式管理冷热数据。
//
// 读路径：先查 Cache → miss 时查 Storage → 回写 Cache
// 写路径：先写 Storage → 删除/更新 Cache（保证一致性）
type MemoryService struct {
	store *SQLiteStorage
	cache *Cache
}

func NewMemoryService(dbPath string) (*MemoryService, error) {
	store, err := NewSQLiteStorage(dbPath)
	if err != nil {
		return nil, err
	}
	return &MemoryService{
		store: store,
		cache: NewCache(),
	}, nil
}

// ── 用户画像 ─────────────────────────────────────────────────────────────────

func (m *MemoryService) GetOrCreateProfile(ctx context.Context, userID string) *UserProfile {
	cacheKey := KeyProfile + userID
	if v := m.cache.Get(cacheKey); v != nil {
		return v.(*UserProfile)
	}
	p, err := m.store.GetUserProfile(ctx, userID)
	if err != nil {
		log.Printf("memory: get profile: %v", err)
	}
	if p == nil {
		p = &UserProfile{
			UserID:       userID,
			OverallLevel: 1,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		m.store.UpsertUserProfile(ctx, p)
	}
	m.cache.Set(cacheKey, p, TTLProfile)
	return p
}

func (m *MemoryService) UpdateProfile(ctx context.Context, p *UserProfile) {
	p.UpdatedAt = time.Now()
	m.store.UpsertUserProfile(ctx, p)
	m.cache.Set(KeyProfile+p.UserID, p, TTLProfile)
}

// ── 薄弱点管理 ───────────────────────────────────────────────────────────────

// GetWeaknesses Cache-Aside 读取薄弱点列表。
func (m *MemoryService) GetWeaknesses(ctx context.Context, userID string) []*WeaknessRecord {
	cacheKey := KeyWeakness + userID
	if v := m.cache.Get(cacheKey); v != nil {
		return v.([]*WeaknessRecord)
	}
	list, err := m.store.GetWeaknesses(ctx, userID)
	if err != nil {
		log.Printf("memory: get weakness: %v", err)
		return nil
	}
	m.cache.Set(cacheKey, list, TTLWeakness)
	return list
}

// RecordAnswer 记录一次答题，自动更新薄弱点。
//
// 规则：
//   - score < 60 → 记录/加重薄弱点，OccurrenceCount+1
//   - score > 85 → 减轻薄弱点；若 OccurrenceCount=0 则直接移除
//   - 60-85 → 无操作（中性区间）
func (m *MemoryService) RecordAnswer(ctx context.Context, r *QuestionRecord) {
	r.CreatedAt = time.Now()
	// 写 DB
	if err := m.store.AddQuestionRecord(ctx, r); err != nil {
		log.Printf("memory: add record: %v", err)
	}
	// 失效 recent cache
	m.cache.Del(KeyRecentRec + r.UserID)

	// 更新薄弱点
	for _, tag := range r.Tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}
		m.updateWeakness(ctx, r.UserID, tag, r.Score)
	}

	// 更新用户画像
	p := m.GetOrCreateProfile(ctx, r.UserID)
	p.TotalQuestions++
	p.AvgScore = (p.AvgScore*float64(p.TotalQuestions-1) + r.Score) / float64(p.TotalQuestions)
	// 综合等级：近 20 题均分映射到 1-5
	p.OverallLevel = scoreToLevel(p.AvgScore)
	m.UpdateProfile(ctx, p)
}

func (m *MemoryService) updateWeakness(ctx context.Context, userID, tag string, score float64) {
	weaknesses := m.GetWeaknesses(ctx, userID)
	var found *WeaknessRecord
	for _, w := range weaknesses {
		if w.Tag == tag {
			found = w
			break
		}
	}

	switch {
	case score < WeaknessThresholdAdd:
		// 低分 → 记录/加重薄弱点
		if found == nil {
			found = &WeaknessRecord{
				UserID:        userID,
				Tag:           tag,
				OccurrenceCount: 0,
			}
		}
		// EMA 更新薄弱度：新弱点权重更大
		alpha := 0.4
		found.WeaknessScore = alpha*(100-score) + (1-alpha)*found.WeaknessScore
		found.OccurrenceCount++
		found.LastSeen = time.Now()
		found.ExpiresAt = time.Now().AddDate(0, 0, WeaknessExpireDays)
		m.store.UpsertWeakness(ctx, found)

	case score > WeaknessThresholdRemove && found != nil:
		// 高分 → 减轻薄弱点
		found.OccurrenceCount--
		if found.OccurrenceCount <= 0 {
			m.store.RemoveWeakness(ctx, userID, tag)
		} else {
			alpha := 0.3
			found.WeaknessScore = found.WeaknessScore * (1 - alpha)
			found.ExpiresAt = time.Now().AddDate(0, 0, WeaknessExpireDays)
			m.store.UpsertWeakness(ctx, found)
		}
	}

	// 失效 weakness cache
	m.cache.Del(KeyWeakness + userID)
}

// GetTopWeakTags 返回用户最薄弱的 N 个技术标签，用于针对性出题。
func (m *MemoryService) GetTopWeakTags(ctx context.Context, userID string, n int) []string {
	list := m.GetWeaknesses(ctx, userID)
	var tags []string
	for i, w := range list {
		if i >= n {
			break
		}
		tags = append(tags, w.Tag)
	}
	return tags
}

// GetSessionContext 构建当前会话的短期记忆（滑动窗口）。
func (m *MemoryService) GetSessionContext(ctx context.Context, userID, sessionID string) *SessionContext {
	sc := &SessionContext{
		SessionID: sessionID,
		UserID:    userID,
		MaxWindow: ContextWindowSize,
	}
	// 从最近记录中恢复上下文
	records, err := m.store.GetRecentRecords(ctx, userID, ContextWindowSize)
	if err != nil {
		return sc
	}
	for _, r := range records {
		sc.Window = append(sc.Window, ContextMessage{
			Role:    "interviewer",
			Content: r.Question,
			Phase:   r.Phase,
			Score:   r.Score,
		})
		if r.Answer != "" {
			sc.Window = append(sc.Window, ContextMessage{
				Role:    "candidate",
				Content: r.Answer,
				Phase:   r.Phase,
				Score:   r.Score,
			})
		}
	}
	return sc
}

// PruneExpired 清理过期薄弱点（建议定期调用）。
func (m *MemoryService) PruneExpired(ctx context.Context) {
	if err := m.store.PruneExpiredWeaknesses(ctx); err != nil {
		log.Printf("memory: prune: %v", err)
	}
}

// BuildWeaknessPrompt 把薄弱点信息拼成 LLM 可用的 Prompt 片段。
func (m *MemoryService) BuildWeaknessPrompt(ctx context.Context, userID string) string {
	tags := m.GetTopWeakTags(ctx, userID, 5)
	if len(tags) == 0 {
		return ""
	}
	return fmt.Sprintf(
		"\n\n【候选人薄弱点（请重点考察）】：%s\n请结合上述薄弱点设计追问，帮助候选人发现不足。",
		strings.Join(tags, "、"),
	)
}

func scoreToLevel(avg float64) int {
	switch {
	case avg >= 85:
		return 5
	case avg >= 75:
		return 4
	case avg >= 65:
		return 3
	case avg >= 55:
		return 2
	default:
		return 1
	}
}
