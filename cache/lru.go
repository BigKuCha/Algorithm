package cache

import (
	"container/list"
	"errors"
)

type Node struct {
	Key, Value interface{}
}

// LRU是Least Recently Used的缩写，即最近最少使用
// 当缓存容量满了，需要淘汰最近最少使用的的数据
// 使用双链表存储数据，存储后返回节点地址作为map的value
type LRUCache struct {
	Cap      int // 容量
	dlist    *list.List
	cacheMap map[interface{}]*list.Element
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		Cap:      cap,
		dlist:    list.New(),
		cacheMap: make(map[interface{}]*list.Element),
	}
}

// 缓存容量
func (l *LRUCache) Len() int {
	return l.dlist.Len()
}

func (l *LRUCache) Set(k, v interface{}) error {
	if l.dlist == nil {
		return errors.New("LRUCache未初始化")
	}
	// 如果存在key，链表位置移动到首位
	if pElement, ok := l.cacheMap[k]; ok {
		l.dlist.MoveToFront(pElement)
		pElement.Value.(*Node).Value = v
		return nil
	}
	newNode := &Node{
		Key:   k,
		Value: v,
	}
	newElement := l.dlist.PushFront(newNode)
	l.cacheMap[k] = newElement
	if l.Len() > l.Cap {
		lastElement := l.dlist.Back()
		if lastElement == nil {
			return nil
		}
		delete(l.cacheMap, lastElement.Value.(*Node).Key)
		l.dlist.Remove(lastElement)
	}

	return nil
}

func (l *LRUCache) Get(k interface{}) (v interface{}, exist bool, error error) {
	if l.dlist == nil {
		return v, false, errors.New("LRUCache 未初始化")
	}
	if element, ok := l.cacheMap[k]; ok {
		l.dlist.MoveToFront(element)
		return element.Value.(*Node).Value, true, nil
	}
	return v, false, nil
}

func (l *LRUCache) Remove(k interface{}) bool {
	if l.dlist == nil {
		return false
	}
	if element, ok := l.cacheMap[k]; ok {
		l.dlist.Remove(element)
		delete(l.cacheMap, k)
		return true
	}
	return false
}
