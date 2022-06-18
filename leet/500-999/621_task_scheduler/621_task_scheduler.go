package _621_task_scheduler

import "container/heap"

// tptl. passed
func leastInterval(tasks []byte, n int) int {
	freq := make(map[byte]int)
	for _, task := range tasks {
		freq[task]++
	}

	h := maxHeap{freq: freq}
	for task := range freq {
		heap.Push(&h, task)
	}

	var waitList []byte
	res := 0
	for h.Len() > 0 {
		waitList = waitList[:0]
		j := n + 1
		for ; j > 0 && h.Len() > 0; j-- {
			res++
			curr := heap.Pop(&h).(byte)
			if freq[curr] > 1 {
				freq[curr]--
				waitList = append(waitList, curr)
			}
		}

		for _, task := range waitList {
			heap.Push(&h, task)
		}

		if h.Len() > 0 {
			res += j // idle
		}
	}

	return res
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
