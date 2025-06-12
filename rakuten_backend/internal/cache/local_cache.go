package cache

import (
	"time"

	"github.com/dgraph-io/ristretto"
	"log"
)

type CacheItem[T any] struct {
	Data      T
	ExpiresAt int64
}

type LocalCache[T any] struct {
	client *ristretto.Cache
}

func NewLocalCache[T any](maxMemoryMB int64) *LocalCache[T] {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 10_000_000,        // 推荐为 MaxCost 的 10 倍
		MaxCost:     maxMemoryMB << 20, // 内存限制，单位：字节
		BufferItems: 64,
	})
	if err != nil {
		log.Fatalf("Failed to create cache: %v", err)
	}

	return &LocalCache[T]{client: cache}
}

// Set 设置缓存，ttl 单位为秒
func (c *LocalCache[T]) Set(key string, data T, ttlSeconds int64) {
	item := CacheItem[T]{
		Data:      data,
		ExpiresAt: time.Now().Unix() + ttlSeconds,
	}
	c.client.Set(key, item, 1) // cost 可以根据数据大小调整
	c.client.Wait()
}

// Get 获取缓存
func (c *LocalCache[T]) Get(key string) (T, bool) {
	var zero T
	val, found := c.client.Get(key)
	if !found {
		return zero, false
	}

	item, ok := val.(CacheItem[T])
	if !ok {
		return zero, false
	}

	if time.Now().Unix() > item.ExpiresAt {
		c.client.Del(key)
		return zero, false
	}

	return item.Data, true
}

// Del 删除
func (c *LocalCache[T]) Del(key string) {
	c.client.Del(key)
}
