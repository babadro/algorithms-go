package _2402_meeting_rooms_iii

import (
	"container/heap"
	"sort"
)

var curTime int

// #bnsrg
// tptl passed, but slow
// todo 1 refac - make it faster and shorter
func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	rooms := make(minHeap, n)
	for i := range rooms {
		rooms[i].id = i
	}

	for _, m := range meetings {
		curTime = m[0]

		heap.Init(&rooms)

		r := heap.Pop(&rooms).(room)

		r.counter++

		if r.freeFrom > m[0] {
			r.freeFrom += m[1] - m[0]
		} else {
			r.freeFrom = m[1]
		}

		heap.Push(&rooms, r)
	}

	resHeap := resMinHeap(rooms)

	heap.Init(&resHeap)

	return resHeap[0].id
}

type room struct {
	id       int
	freeFrom int
	counter  int
}

type minHeap []room

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	isFreeI, isFreeJ := h[i].freeFrom <= curTime, h[j].freeFrom <= curTime
	idI, idJ := h[i].id, h[j].id
	freeFromI, freeFromJ := h[i].freeFrom, h[j].freeFrom

	if isFreeI && isFreeJ {
		return idI < idJ
	}

	if !isFreeI && !isFreeJ {
		if freeFromI != freeFromJ {
			return freeFromI < freeFromJ
		}

		return idI < idJ
	}

	if isFreeI {
		return true
	}

	return false
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.(room))
}

func (h *minHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}

type resMinHeap []room

func (h resMinHeap) Len() int {
	return len(h)
}

func (h resMinHeap) Less(i, j int) bool {
	counterI, counterJ := h[i].counter, h[j].counter

	if counterI != counterJ {
		return counterI > counterJ
	}

	return h[i].id < h[j].id
}

func (h resMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *resMinHeap) Push(v interface{}) {
	*h = append(*h, v.(room))
}

func (h *resMinHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}
