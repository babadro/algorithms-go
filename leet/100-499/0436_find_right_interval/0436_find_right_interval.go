package _436_find_right_interval

import (
	"container/heap"
	"sort"
)

// tptl. passed. fast
func findRightInterval2(intervals [][]int) []int {
	starts, ends := make([]int, len(intervals)), make([]int, len(intervals))
	for i := range intervals {
		starts[i], ends[i] = i, i
	}

	sort.Slice(starts, func(i, j int) bool {
		return intervals[starts[i]][0] < intervals[starts[j]][0]
	})

	sort.Slice(ends, func(i, j int) bool {
		return intervals[ends[i]][1] < intervals[ends[j]][1]
	})

	res := make([]int, len(intervals))

	j := len(ends) - 1
	for i := len(ends) - 1; i >= 0; i-- {
		topEnd := ends[i]
		topStart := -1
		for ; j >= 0 && intervals[starts[j]][0] >= intervals[topEnd][1]; j-- {
			topStart = starts[j]
		}

		res[topEnd] = topStart
		if topStart != -1 {
			j++
		}
	}

	return res
}

// two heaps solution. passed but slow
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
