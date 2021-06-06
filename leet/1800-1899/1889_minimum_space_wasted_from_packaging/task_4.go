package _1889_minimum_space_wasted_from_packaging

import (
	"container/heap"
	"sort"
)

// todo 1 it works, but slow - 4188ms. Find faster solution
func minWastedSpace(packages []int, boxes [][]int) int {
	sort.Ints(packages)
	n := len(packages)
	min, max := packages[0], packages[n-1]

	res := -1
Loop:
	for _, line := range boxes {
		flag := false
		h := &minHeap{}
		for _, box := range line {
			if box >= max {
				flag = true
			}
		}

		if !flag {
			continue
		}

		for _, box := range line {
			if box >= min {
				heap.Push(h, box)
			}
		}

		currRes := 0
		currBox := heap.Pop(h).(int)

		for _, pack := range packages {
			for currBox < pack {
				currBox = heap.Pop(h).(int)
			}

			currRes += currBox - pack
			if res != -1 && currRes >= res {
				continue Loop
			}
		}

		if res == -1 || currRes < res {
			res = currRes
		}

		if res == 0 {
			return 0
		}

	}

	return res % 1_000_000_007
}

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]

	return res
}
