package _1_top_k_numbers

import "container/heap"

func findKLargestNumbers(nums []int, k int) []int {
	mh := minHeap{}
	for i := 0; i < k; i++ {
		heap.Push(&mh, nums[i])
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > mh[0] {
			heap.Pop(&mh)
			heap.Push(&mh, nums[i])
		}
	}

	return mh
}

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *minHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}
