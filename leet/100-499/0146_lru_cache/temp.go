package _146_lru_cache

import "container/heap"

type item3 struct {
	key, value, timestamp int
}

type LRUCache3 struct {
	h         minHeap
	capacity  int
	timestamp int
}

func Constructor3(capacity int) LRUCache3 {
	keyToIdx := make(map[int]int, capacity)
	return LRUCache3{
		capacity:  capacity,
		timestamp: 0,
		h: minHeap{
			keyToIDx: keyToIdx,
			items:    make([]item3, 0, capacity),
		},
	}
}

func (this *LRUCache3) Get(key int) int {
	idx, ok := this.h.keyToIDx[key]
	if !ok {
		return -1
	}

	this.timestamp++

	cacheItem := this.h.items[idx]
	cacheItem.timestamp = this.timestamp
	heap.Fix(&this.h, idx)

	return cacheItem.value
}

func (this *LRUCache3) Put(key int, value int) {
	idx, ok := this.h.keyToIDx[key]
	this.timestamp++

	if ok {
		this.h.items[idx].value, this.h.items[idx].timestamp = value, this.timestamp
		heap.Fix(&this.h, idx)

		return
	}

	heap.Push(&this.h, item3{key: key, value: value, timestamp: this.timestamp})
	if this.h.Len() > this.capacity {
		_ = heap.Pop(&this.h)
	}
}

type minHeap struct {
	keyToIDx map[int]int
	items    []item3
}

func (h minHeap) Len() int { return len(h.items) }
func (h minHeap) Less(i, j int) bool {
	return h.items[i].timestamp < h.items[j].timestamp
}
func (h minHeap) Swap(i, j int) {
	h.keyToIDx[h.items[i].key] = j
	h.keyToIDx[h.items[j].key] = i

	h.items[i], h.items[j] = h.items[j], h.items[i]
}
func (h *minHeap) Push(v interface{}) {
	cacheItem := v.(item3)
	h.keyToIDx[cacheItem.key] = len(h.items)
	h.items = append(h.items, cacheItem)
}
func (h *minHeap) Pop() interface{} {
	last := len(h.items) - 1
	res := h.items[last]
	h.items = h.items[:last]

	delete(h.keyToIDx, res.key)

	return res
}
