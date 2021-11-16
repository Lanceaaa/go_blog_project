package cache

import (
	"log"
	"sync"
)

// 设置/添加一个缓存，如果 key 存在，用新值覆盖旧值；
// 通过 key 获取一个缓存值；
// 通过 key 删除一个缓存值；
// 删除最“无用”的一个缓存值；
// 获取缓存已存在的记录数；
type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Del(key string)
	DelOldest()
	Len() int
}

const DefaultMaxBytes = 1 << 29

type safeCache struct {
	m     sync.RWMutex
	cache Cache

	nhit, nget int
}

func newSafeCache(cache Cache) *safeCache {
	return &safeCache{
		cache: cache,
	}
}

func (sc *safeCache) set(key string, value interface{}) {
	sc.m.Lock()
	defer sc.m.Unlock()
	sc.cache.Set(key, value)
}

func (sc *safeCache) get(key string) interface{} {
	sc.m.RLock()
	defer sc.m.RUnlock()
	sc.nget++
	if sc.cache == nil {
		return nil
	}

	v := sc.cache.Get(key)
	if v != nil {
		log.Println("[TourCache] hit")
		sc.nhit++
	}

	return v
}

type Stat struct {
	NHit, NGet int
}

func (sc *safeCache) stat() *Stat {
	sc.m.RLock()
	defer sc.m.RUnlock()
	return &Stat{
		NHit: sc.nhit,
		NGet: sc.nget,
	}
}
