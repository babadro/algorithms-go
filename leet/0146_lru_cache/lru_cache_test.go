package _146_lru_cache

import "testing"

func TestLruCache(t *testing.T) {
	cache := Constructor2(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	t.Log(cache.Get(1)) // returns 1
	cache.Put(3, 3)     // evicts key 2
	t.Log(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)     // evicts key 1
	t.Log(cache.Get(1)) // returns -1 (not found)
	t.Log(cache.Get(3)) // returns 3
}
func TestLruCache2(t *testing.T) {
	cache := Constructor2(1)
	cache.Put(1, 1)
	cache.Put(2, 2)
	t.Log(cache.Get(1)) // returns -1
	cache.Put(3, 3)
	t.Log(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)
	t.Log(cache.Get(1)) // returns -1 (not found)
	t.Log(cache.Get(3)) // returns -1
	t.Log(cache.Get(4)) // returns 4
}

func TestLruCache3(t *testing.T) {
	cache := Constructor2(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	t.Log(cache.Get(2)) // returns 2
	cache.Put(1, 1)
	cache.Put(4, 1)
	t.Log(cache.Get(2)) // returns -1 (not found)
}

func TestLruCache4(t *testing.T) {
	cache := Constructor2(2)
	t.Log(cache.Get(2)) // -1
	cache.Put(2, 6)
	t.Log(cache.Get(1)) // -1
	cache.Put(1, 5)
	cache.Put(1, 2)
	t.Log(cache.Get(1)) // 2
	t.Log(cache.Get(2)) // 6
}

// TODO 1
//Input
//["LRUCache2","put","put","get","put","put","get"]
//[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
//Output
//[null,null,null,1,null,null,-1]
//Expected
//[null,null,null,2,null,null,-1]
