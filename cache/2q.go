package cache

type TwoQueuesCache struct {
	recent *LRUCache
	cache  *LRUCache
}

func NewTwoQueuesCache(recentCap, cacheCap int) *TwoQueuesCache {
	recent := NewLRUCache(recentCap)
	cache := NewLRUCache(cacheCap)
	return &TwoQueuesCache{
		recent: recent,
		cache:  cache,
	}
}

func (t *TwoQueuesCache) Has(k interface{}) bool {
	return t.recent.Has(k) || t.cache.Has(k)
}

func (t *TwoQueuesCache) Len() int {
	return t.recent.Len() + t.cache.Len()
}

func (t *TwoQueuesCache) Set(k, v interface{}) {
	// 缓存中有值，直接修改
	if t.cache.Has(k) {
		t.cache.Set(k, v)
		return
	}

	// 最近访问列表里有，则移到缓存列表，同时删除最近访问列表
	if t.recent.Has(k) {
		t.cache.Set(k, v)
		t.recent.Remove(k)
		return
	}
	// 加入最近访问列表
	t.recent.Set(k, v)
}

func (t *TwoQueuesCache) Get(k interface{}) (interface{}, error) {
	if t.cache.Has(k) {
		return t.cache.Get(k)
	}
	if t.recent.Has(k) {
		return t.recent.Get(k)
	}
	return nil, nil
}

func (t *TwoQueuesCache) Remove(k interface{}) bool {
	if t.recent.Remove(k) {
		return true
	}
	if t.cache.Remove(k) {
		return true
	}
	return false
}
