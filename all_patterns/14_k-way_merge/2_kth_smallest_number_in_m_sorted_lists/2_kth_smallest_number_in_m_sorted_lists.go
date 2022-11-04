package __kth_smallest_number_in_m_sorted_lists

import (
	"container/heap"
	"math"
)

// leetcode 378
func findKthSmallest(lists [][]int, k int) int {
	numIDxes := make([]int, len(lists))
	h := make(arrItems, 0, len(lists))
	for arrIDx := range numIDxes {
		arr := lists[arrIDx]
		if len(arr) > 0 {
			h = append(h, [2]int{arr[0], arrIDx})
			numIDxes[arrIDx] = 1
		}
	}

	if len(h) == 0 {
		return math.MaxInt64
	}
	heap.Init(&h)

	for i := 0; len(h) > 0; i++ {
		minItem := heap.Pop(&h).([2]int)
		num, arrIDx := minItem[0], minItem[1]

		if i+1 == k {
			return num
		}

		if numIDx := numIDxes[arrIDx]; numIDx < len(lists[arrIDx]) {
			heap.Push(&h, [2]int{lists[arrIDx][numIDx], arrIDx})
			numIDxes[arrIDx]++
		}
	}

	return math.MaxInt64
}

type arrItems [][2]int

func (a arrItems) Len() int {
	return len(a)
}
func (a arrItems) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}
func (a arrItems) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a *arrItems) Push(v interface{}) {
	*a = append(*a, v.([2]int))
}
func (a *arrItems) Pop() interface{} {
	last := len(*a) - 1
	res := (*a)[last]
	*a = (*a)[:last]

	return res
}
