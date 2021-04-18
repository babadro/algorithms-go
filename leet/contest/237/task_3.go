package _37

import (
	"container/heap"
	"math"
	"sort"
)

// todo 1 doesn't work
func getOrder(tasks [][]int) []int {
	n := len(tasks)

	newTasks := make([][]int, n)
	for i, t := range tasks {
		newTasks[i] = []int{t[0], t[1], i}
	}

	sort.Slice(newTasks, func(i, j int) bool {
		return newTasks[i][0] < newTasks[j][0]
	})

	res := make([]int, 0, n)

	taskFinishTime := 1
	h := &minHeap{}
	var currentTime int
	for i := 0; i <= n; {
		if i == n {
			currentTime = math.MaxInt32
		} else {
			currentTime = newTasks[i][0]
		}

		for taskFinishTime < currentTime && len(*h) > 0 {
			taskFromQueue := heap.Pop(h).(task)
			taskFinishTime += taskFromQueue.processingTime
			res = append(res, taskFromQueue.idx)
		}

		if i == n {
			break
		}

		for i < n && newTasks[i][0] == currentTime {
			heap.Push(h, task{idx: newTasks[i][2], processingTime: newTasks[i][1]})
			i++
		}
	}

	return res
}

type task struct {
	idx            int
	processingTime int
}

type minHeap []task

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].processingTime < h[j].processingTime }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(task)) }
func (h *minHeap) Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}
