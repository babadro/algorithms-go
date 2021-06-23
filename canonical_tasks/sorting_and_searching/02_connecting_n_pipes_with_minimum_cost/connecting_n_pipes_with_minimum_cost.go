package _2_connecting_n_pipes_with_minimum_cost

import "container/heap"

// leetcode 1167 (premium)
// passed, I think. medium. tptl
func connectSticks(sticks []int) int {
	h := minHeap(sticks)

	heap.Init(&h)

	var cost int

	for h.Len() >= 2 {
		curr := heap.Pop(&h).(int) + heap.Pop(&h).(int)
		cost += curr
		heap.Push(&h, curr)
	}

	return cost
}

type minHeap []int

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Len() int {
	return len(h)
}
func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *minHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}
