package _767_reorganize_string

import "container/heap"

// tptl. passed. todo2 find faster solution if exists
func reorganizeString(s string) string {
	freq := make(map[byte]int, 26)
	for i := range s {
		freq[s[i]]++
	}

	var arr []byte
	for char := range freq {
		arr = append(arr, char)
	}

	h := maxHeap{freq: freq, arr: arr}
	heap.Init(&h)

	res := make([]byte, 0, len(s))
	var prev byte
	for i := 0; h.Len() > 0; i++ {
		cur := heap.Pop(&h).(byte)
		if i > 0 && freq[prev] > 0 {
			heap.Push(&h, prev)
		}

		res = append(res, cur)
		freq[cur]--

		prev = cur
	}

	if len(res) == len(s) {
		return string(res)
	}

	return ""
}

type maxHeap struct {
	freq map[byte]int
	arr  []byte
}

func (h maxHeap) Len() int            { return len(h.arr) }
func (h maxHeap) Less(i, j int) bool  { return h.freq[h.arr[i]] > h.freq[h.arr[j]] }
func (h maxHeap) Swap(i, j int)       { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }
func (h *maxHeap) Push(v interface{}) { (*h).arr = append((*h).arr, v.(byte)) }
func (h *maxHeap) Pop() interface{} {
	last := len(h.arr) - 1
	res := h.arr[last]
	(*h).arr = (*h).arr[:last]

	return res
}
