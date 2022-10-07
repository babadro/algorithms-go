package _373_find_k_pairs_with_smalles_sums

import "container/heap"

// tptl. passed. medium (hard for me)
// todo 2 find solution without heap - it should be faster
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var h maxHeap

	for i := 0; i < len(nums1) && i < k; i++ {
		for j := 0; j < len(nums2) && j < k; j++ {
			num1, num2 := nums1[i], nums2[j]
			if h.Len() < k {
				heap.Push(&h, []int{num1, num2})
			} else {
				top := h[0]
				if num1+num2 >= top[0]+top[1] {
					break
				}

				heap.Pop(&h)
				heap.Push(&h, []int{num1, num2})
			}
		}
	}

	return h
}

type maxHeap [][]int

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	return h[i][0]+h[i][1] > h[j][0]+h[j][1]
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *maxHeap) Push(v interface{}) {
	*h = append(*h, v.([]int))
}
func (h *maxHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}
