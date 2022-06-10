package _436_find_right_interval

import (
	"container/heap"
)

// tptl. passed.
// todo 2 it's slow. find faster solution
func findRightInterval(intervals [][]int) []int {
	startH, endH := startMaxHeap{intervals: intervals}, endMaxHeap{intervals: intervals}
	for i := range intervals {
		heap.Push(&startH, i)
		heap.Push(&endH, i)
	}

	res := make([]int, len(intervals))
	for _ = range intervals {
		topEnd := heap.Pop(&endH).(int)
		topStart := -1
		for startH.Len() > 0 && intervals[startH.indexes[0]][0] >= intervals[topEnd][1] {
			topStart = heap.Pop(&startH).(int)
		}
		res[topEnd] = topStart
		if topStart != -1 {
			heap.Push(&startH, topStart)
		}
	}

	return res
}

type startMaxHeap struct {
	intervals [][]int
	indexes   []int
}

func (h startMaxHeap) Len() int { return len(h.indexes) }
func (h startMaxHeap) Less(i, j int) bool {
	return h.intervals[h.indexes[i]][0] > h.intervals[h.indexes[j]][0]
}
func (h startMaxHeap) Swap(i, j int)       { h.indexes[i], h.indexes[j] = h.indexes[j], h.indexes[i] }
func (h *startMaxHeap) Push(v interface{}) { (*h).indexes = append((*h).indexes, v.(int)) }
func (h *startMaxHeap) Pop() interface{} {
	last := len(h.indexes) - 1
	res := h.indexes[last]
	(*h).indexes = (*h).indexes[:last]

	return res
}

type endMaxHeap struct {
	intervals [][]int
	indexes   []int
}

func (h endMaxHeap) Len() int { return len(h.indexes) }
func (h endMaxHeap) Less(i, j int) bool {
	return h.intervals[h.indexes[i]][1] > h.intervals[h.indexes[j]][1]
}
func (h endMaxHeap) Swap(i, j int)       { h.indexes[i], h.indexes[j] = h.indexes[j], h.indexes[i] }
func (h *endMaxHeap) Push(v interface{}) { (*h).indexes = append((*h).indexes, v.(int)) }
func (h *endMaxHeap) Pop() interface{} {
	last := len(h.indexes) - 1
	res := h.indexes[last]
	(*h).indexes = (*h).indexes[:last]

	return res
}
