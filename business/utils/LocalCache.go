package utils

import (
	"sync"
	"time"
)

type CacheItem struct {
	value      string
	expiration time.Time
}

type LocalCache struct {
	cache map[string]CacheItem
	mu    sync.RWMutex
}

func NewLocalCache() *LocalCache {
	return &LocalCache{
		cache: make(map[string]CacheItem),
	}
}

func (lc *LocalCache) Delete(key string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.cache, key)
}

func (lc *LocalCache) Set(key string, value string, expiration time.Duration) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.cache[key] = CacheItem{
		value:      value,
		expiration: time.Now().Add(expiration),
	}
}

func (lc *LocalCache) Get(key string) (string, bool) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	item, found := lc.cache[key]
	if !found {
		return "", false
	}

	if time.Now().After(item.expiration) {
		// 如果缓存已过期，则删除该项并返回未找到
		delete(lc.cache, key)
		return "", false
	}

	return item.value, true
}
