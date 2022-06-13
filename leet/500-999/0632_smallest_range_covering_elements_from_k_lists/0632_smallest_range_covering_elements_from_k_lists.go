package _632_smallest_range_covering_elements_from_k_lists

import (
	"container/heap"
	"math"
)

// tptl passed
func smallestRange(nums [][]int) []int {
	h := minHeap{nums: nums}
	currMax := math.MinInt64
	for arrIdx := range nums {
		heap.Push(&h, [2]int{arrIdx, 0})
		currMax = max(currMax, nums[arrIdx][0])
	}

	start, end := 0, math.MaxInt64
	for h.Len() == len(nums) {
		minEl := heap.Pop(&h).([2]int)
		arrIdx, numIdx := minEl[0], minEl[1]
		num := nums[arrIdx][numIdx]
		if currMax-num < end-start {
			start, end = num, currMax
		}
		numIdx++
		if len(nums[arrIdx]) > numIdx {
			heap.Push(&h, [2]int{arrIdx, numIdx})
			currMax = max(currMax, nums[arrIdx][numIdx])
		}
	}

	return []int{start, end}
}

type minHeap struct {
	items [][2]int
	nums  [][]int
}

func (h minHeap) Len() int { return len(h.items) }
func (h minHeap) Less(i, j int) bool {
	arrIdxI, numIdxI, arrIdxJ, numIdxJ := h.items[i][0], h.items[i][1], h.items[j][0], h.items[j][1]
	return h.nums[arrIdxI][numIdxI] < h.nums[arrIdxJ][numIdxJ]
}
func (h minHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *minHeap) Push(v interface{}) {
	h.items = append(h.items, v.([2]int))
}

func (h *minHeap) Pop() interface{} {
	last := len(h.items) - 1
	res := h.items[last]
	h.items = h.items[:last]

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
