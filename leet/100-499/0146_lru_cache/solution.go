package _146_lru_cache

import "container/list"

// tptl. passed.
type cacheItem struct {
	key, val int
}

type LRUCache4 struct {
	l        *list.List
	m        map[int]*list.Element
	capacity int
}

func Constructor4(capacity int) LRUCache4 {
	return LRUCache4{
		l:        list.New(),
		m:        make(map[int]*list.Element, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache4) Get(key int) int {
	el, ok := this.m[key]
	if !ok {
		return -1
	}

	this.l.MoveToFront(el)
	return el.Value.(cacheItem).val
}

func (this *LRUCache4) Put(key int, value int) {
	if el, ok := this.m[key]; ok {
		el.Value = cacheItem{
			key: key,
			val: value,
		}
		this.l.MoveToFront(el)

		return
	}

	newEl := this.l.PushFront(cacheItem{
		key: key,
		val: value,
	})
	this.m[key] = newEl

	if this.l.Len() > this.capacity {
		last := this.l.Back()
		this.l.Remove(last)
		delete(this.m, last.Value.(cacheItem).key)
	}
}
