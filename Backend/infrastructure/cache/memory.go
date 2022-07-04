package cache

import "sync"

// MemoryCache 内存缓存器，使用sync.Map作为底层，线程安全
type MemoryCache struct {
	data sync.Map
}

// Set 向缓存器中写入一个key-value对
func (m *MemoryCache) Set(key, value interface{}) {
	m.data.Store(key, value)
}

// Get 向缓存器中获取一个key-value对
func (m *MemoryCache) Get(key interface{}) (value interface{}, ok bool) {
	return m.data.Load(key)
}

// GetOrSet 向缓存器中读取，没有则写入一个key-value对
func (m *MemoryCache) GetOrSet(key, value interface{}) (actual interface{}, ok bool) {
	return m.data.LoadOrStore(key, value)
}

// Remove 从缓存器中删除指定的key-value对
func (m *MemoryCache) Remove(key interface{}) {
	m.data.Delete(key)
}
