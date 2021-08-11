package _962_remove_stones_to_minimize_the_total

import "container/heap"

// passed. easy
func minStoneSum(piles []int, k int) int {
	sum := 0
	for _, num := range piles {
		sum += num
	}

	h := maxHeap(piles)
	heap.Init(&h)

	for k > 0 {
		k--
		count := h[0] / 2
		if count == 0 {
			break
		}

		h[0] -= count
		sum -= count
		heap.Fix(&h, 0)
	}

	return sum
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *maxHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *maxHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}
