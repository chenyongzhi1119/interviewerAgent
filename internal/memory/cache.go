package memory

import (
	"sync"
	"time"
)

// Cache 内存热数据缓存层，模拟 Redis Cache-Aside 模式。
// 生产环境可替换为真实 Redis 客户端，接口保持不变。
type Cache struct {
	mu      sync.RWMutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	value     any
	expiresAt time.Time
}

func NewCache() *Cache {
	c := &Cache{entries: make(map[string]cacheEntry)}
	go c.gcLoop()
	return c
}

// Set 写入缓存，ttl 为 0 则永不过期。
func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	exp := time.Time{}
	if ttl > 0 {
		exp = time.Now().Add(ttl)
	}
	c.entries[key] = cacheEntry{value: value, expiresAt: exp}
}

// Get 读缓存，miss 时返回 nil。
func (c *Cache) Get(key string) any {
	c.mu.RLock()
	defer c.mu.RUnlock()
	e, ok := c.entries[key]
	if !ok {
		return nil
	}
	if !e.expiresAt.IsZero() && time.Now().After(e.expiresAt) {
		return nil // expired
	}
	return e.value
}

// Del 主动失效。
func (c *Cache) Del(key string) {
	c.mu.Lock()
	delete(c.entries, key)
	c.mu.Unlock()
}

// gcLoop 定期清理过期条目。
func (c *Cache) gcLoop() {
	tick := time.NewTicker(5 * time.Minute)
	for range tick.C {
		now := time.Now()
		c.mu.Lock()
		for k, e := range c.entries {
			if !e.expiresAt.IsZero() && now.After(e.expiresAt) {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}

// CacheKey 统一 key 命名。
const (
	KeyProfile    = "profile:"   // + userID
	KeyWeakness   = "weakness:"  // + userID
	KeyRecentRec  = "recent:"    // + userID
	KeyTagScore   = "tagscore:"  // + userID

	TTLProfile   = 10 * time.Minute
	TTLWeakness  = 5 * time.Minute
	TTLRecentRec = 2 * time.Minute
)
