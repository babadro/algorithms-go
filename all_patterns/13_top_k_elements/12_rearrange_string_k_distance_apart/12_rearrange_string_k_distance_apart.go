package _12_rearrange_string_k_distance_apart

import "container/heap"

// leetcode premium https://leetcode.com/problems/rearrange-string-k-distance-apart/
// tptl
func reorganizeString(str string, k int) string {
	if k <= 1 {
		return str
	}

	freq := make(map[byte]int)
	for i := range str {
		freq[str[i]]++
	}

	h := maxHeap{freq: freq}
	for char := range freq {
		heap.Push(&h, char)
	}

	res := make([]byte, 0, len(str))
	var queue []byte
	for h.Len() > 0 {
		cur := heap.Pop(&h).(byte)
		res = append(res, cur)

		freq[cur]--
		queue = append(queue, cur)
		if len(queue) == k {
			char := queue[0]
			queue = queue[1:]
			if freq[char] > 0 {
				heap.Push(&h, char)
			}
		}
	}

	if len(res) == len(str) {
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
