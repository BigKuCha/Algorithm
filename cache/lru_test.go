package cache

import "testing"

func TestLRUCache_Set(t *testing.T) {
	lru := NewLRUCache(3)
	if lru.Len() != 0 {
		t.Error("容量错误", lru.Len())
	}
	//设置缓存
	lru.Set("name", "joe")
	lru.Set("age", 18)
	lru.Set("sex", "male")

	if lru.Len() != 3 {
		t.Error("容量错误", lru.Len())
	}

	// 获取缓存
	name, exist, err := lru.Get("name")
	if !exist {
		t.Error("缓存获取失败", err)
	}
	t.Log("姓名:", name)
	lru.Set("name", "lily")
	name, exist, err = lru.Get("name")
	if !exist {
		t.Error("缓存获取失败", err)
	}
	t.Log("新名字", name)

	lru.Set("address", "beijing")
	if lru.Len() != 3 {
		t.Error("容量错误", lru.Len())
	}
	// 超出容量，第一次设置的sex应当移除
	age, exist, err := lru.Get("age")
	if exist {
		t.Error("超出容量，没有删除最近最少访问的key, sex:", age)
	}

	// 移除缓存
	lru.Remove("age")
	name, exist, err = lru.Get("age")
	if exist {
		t.Error("key已经删除，还能获取到", err)
	}
}
