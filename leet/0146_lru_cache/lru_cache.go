package _146_lru_cache

// Doesn't work

type item struct {
	payload, lastUsage int
}

type LRUCache struct {
	newest       int
	counterToKey map[int]int
	storage      map[int]item
	oldest       int
	size         int
	capacity     int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{counterToKey: make(map[int]int, capacity), storage: make(map[int]item, capacity), capacity: capacity}
}

func (this *LRUCache) Get(key int) int {
	if it, ok := this.storage[key]; ok {
		delete(this.counterToKey, it.lastUsage)
		this.newest++
		this.counterToKey[this.newest] = key
		if this.oldest == it.lastUsage {
			this.findNextOldest()
		}
		this.storage[key] = item{payload: it.payload, lastUsage: this.newest}
		return it.payload
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	flag := false
	if oldValue, ok := this.storage[key]; ok {
		if oldValue.lastUsage == this.oldest {
			delete(this.counterToKey, this.oldest)
			flag = true
		}
	} else if this.size == this.capacity {
		keyToDelete := this.counterToKey[this.oldest]
		delete(this.storage, keyToDelete)
		delete(this.counterToKey, this.oldest)
		flag = true
	} else {
		if this.size == 0 {
			this.oldest = 1
			this.counterToKey[1] = key
		}
		this.size++
	}

	this.newest++
	this.storage[key] = item{payload: value, lastUsage: this.newest}
	this.counterToKey[this.newest] = key
	if flag {
		this.findNextOldest()
	}
}

func (this *LRUCache) findNextOldest() {
	for {
		this.oldest++
		if _, ok := this.counterToKey[this.oldest]; ok {
			break
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,item);
 */
