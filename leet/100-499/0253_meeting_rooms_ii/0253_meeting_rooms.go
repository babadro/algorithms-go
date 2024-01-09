package _253_meeting_rooms_ii

import (
	"container/heap"
	"sort"
)

// #bnsrg medium passed
// todo2 there are solutions with sorting without heap
func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var h minHeap

	var res int

	for _, interval := range intervals {
		for h.Len() > 0 && h[0][1] <= interval[0] {
			heap.Pop(&h)
		}

		heap.Push(&h, interval)

		res = max(res, h.Len())
	}

	return res
}

type minHeap [][]int

func (h minHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]

	*h = (*h)[:last]

	return res
}

func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.([]int))
}
