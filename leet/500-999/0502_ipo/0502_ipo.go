package _502_ipo

import "container/heap"

// tptl. passed hard
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	ids := make([]int, len(profits))
	for i := range ids {
		ids[i] = i
	}

	minHeap := projectHeap{
		less: func(i, j int) bool {
			return capital[i] < capital[j]
		},
		ids: ids,
	}

	heap.Init(&minHeap)

	maxHeap := projectHeap{
		less: func(i, j int) bool {
			return profits[i] > profits[j]
		},
	}

	for i := 0; i < k; i++ {
		for minHeap.Len() > 0 {
			idx := minHeap.ids[0]
			minCapital := capital[idx]
			if minCapital > w {
				break
			}

			heap.Pop(&minHeap)
			heap.Push(&maxHeap, idx)
		}

		if maxHeap.Len() == 0 {
			return w
		}

		w += profits[maxHeap.ids[0]]
		heap.Pop(&maxHeap)
	}

	return w
}

type projectHeap struct {
	less func(i, j int) bool
	ids  []int
}

func (h projectHeap) Less(i, j int) bool {
	return h.less(h.ids[i], h.ids[j])
}

func (h projectHeap) Swap(i, j int) {
	h.ids[i], h.ids[j] = h.ids[j], h.ids[i]
}

func (h projectHeap) Len() int {
	return len(h.ids)
}

func (h *projectHeap) Push(v interface{}) {
	h.ids = append(h.ids, v.(int))
}

func (h *projectHeap) Pop() interface{} {
	last := len(h.ids) - 1
	res := h.ids[last]
	h.ids = h.ids[:last]

	return res
}
