package _37

import "container/heap"

// todo 1 doesn't work
func getOrder2(tasks [][]int) []int {
	n := len(tasks)
	res := make([]int, 0, n)

	processingH := &processingTimeHeap{
		tasks: tasks,
		idxes: make([]int, 0, n),
	}

	timeH := &enqueueTimeHeap{
		tasks: tasks,
		idxes: make([]int, 0, n),
	}

	for idx := range tasks {
		heap.Push(timeH, idx)
	}

	currentTime := 1
	for processingH.Len() > 0 || timeH.Len() > 0 {
		if processingH.Len() > 0 {
			idx := heap.Pop(processingH).(int)
			currentTime += tasks[idx][1]

			res = append(res, idx)
		}

		for timeH.Len() > 0 && tasks[timeH.idxes[0]][0] <= currentTime {
			nextIdx := heap.Pop(timeH).(int)

			heap.Push(processingH, nextIdx)
		}

		if processingH.Len() == 0 && timeH.Len() > 0 {
			t := tasks[timeH.idxes[0]][0]
			for timeH.Len() > 0 && tasks[timeH.idxes[0]][0] == t {
				heap.Push(processingH, heap.Pop(timeH).(int))
			}
		}
	}

	return res
}

type processingTimeHeap struct {
	tasks [][]int
	idxes []int
}

func (h processingTimeHeap) Len() int { return len(h.idxes) }
func (h processingTimeHeap) Less(i, j int) bool {
	task1, task2 := h.tasks[h.idxes[i]], h.tasks[h.idxes[j]]

	return task1[1] < task2[1]
}
func (h processingTimeHeap) Swap(i, j int)       { h.idxes[i], h.idxes[j] = h.idxes[j], h.idxes[i] }
func (h *processingTimeHeap) Push(x interface{}) { h.idxes = append(h.idxes, x.(int)) }
func (h *processingTimeHeap) Pop() (v interface{}) {
	idx := h.idxes[len(h.idxes)-1]
	h.idxes = h.idxes[:len(h.idxes)-1]
	return idx
}

type enqueueTimeHeap struct {
	tasks [][]int
	idxes []int
}

func (h enqueueTimeHeap) Len() int            { return len(h.idxes) }
func (h enqueueTimeHeap) Less(i, j int) bool  { return h.tasks[h.idxes[i]][0] < h.tasks[h.idxes[j]][0] }
func (h enqueueTimeHeap) Swap(i, j int)       { h.idxes[i], h.idxes[j] = h.idxes[j], h.idxes[i] }
func (h *enqueueTimeHeap) Push(x interface{}) { h.idxes = append(h.idxes, x.(int)) }
func (h *enqueueTimeHeap) Pop() (v interface{}) {
	idx := h.idxes[len(h.idxes)-1]
	h.idxes = h.idxes[:len(h.idxes)-1]
	return idx
}
