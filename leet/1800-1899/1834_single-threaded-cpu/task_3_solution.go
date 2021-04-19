package _834_single_threaded_cpu

import (
	"container/heap"
	"sort"
)

func getOrder(tasks [][]int) []int {
	n := len(tasks)
	sortedList := make([]task, n)

	for i, t := range tasks {
		sortedList[i] = task{
			startTime:      t[0],
			processingTime: t[1],
			idx:            i,
		}
	}

	sort.Slice(sortedList, func(i, j int) bool {
		return sortedList[i].startTime < sortedList[j].startTime
	})

	minHeap := &priorityQ{}

	res := make([]int, 0, n)
	index, resultIndex, currentTime := 0, 0, 1

	for resultIndex < len(res) {
		for index < len(sortedList) && sortedList[index].startTime <= currentTime {
			heap.Push(minHeap, sortedList[index])
			index++
		}

		if minHeap.Len() == 0 {
			currentTime = sortedList[index].startTime
			continue
		}

		t := heap.Pop(minHeap).(task)
		currentTime += t.processingTime
		res[resultIndex] = t.idx
		resultIndex++
	}

	return res
}

type task struct {
	startTime      int
	processingTime int
	idx            int
}

type priorityQ []task

func (h priorityQ) Len() int { return len(h) }
func (h priorityQ) Less(i, j int) bool {
	if h[i].processingTime == h[j].processingTime {
		return h[i].idx < h[j].idx
	}

	return h[i].idx < h[j].idx
}
func (h priorityQ) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *priorityQ) Push(x interface{}) { *h = append(*h, x.(task)) }
func (h *priorityQ) Pop() (v interface{}) {
	v = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}
