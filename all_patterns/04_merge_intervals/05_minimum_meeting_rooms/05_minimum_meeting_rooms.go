package _5_minimum_meeting_rooms

import (
	"container/heap"
	"sort"

	"github.com/babadro/algorithms-go/utils"
)

// leetcode premium https://leetcode.com/problems/meeting-rooms/
// https://leetcode.com/problems/meeting-rooms-ii/
// tptl
func findMinimumMeetingRooms(meetings [][]int) int {
	if len(meetings) == 0 {
		return 0
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][1]
	})

	minRooms := 0
	h := minHeap{}
	for _, meeting := range meetings {
		for h.Len() > 0 && meeting[0] >= h[0][1] {
			heap.Pop(&h)
		}

		heap.Push(&h, meeting)

		minRooms = utils.Max(minRooms, h.Len())
	}

	return minRooms
}

type minHeap [][]int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i][1] < h[j][1] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *minHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}
