package _2_kth_smallest_number

import "container/heap"

func findKthSmallestNumber(nums []int, k int) int {
	h := maxHeap{}
	heap.Init(&h)

	for i := 0; i < k; i++ {
		heap.Push(&h, nums[i])
	}

	for i := k; i < len(nums); i++ {
		if nums[i] < h[0] {
			heap.Pop(&h)
			heap.Push(&h, nums[i])
		}
	}

	return h[0]
}

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *maxHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}
