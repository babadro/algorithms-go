package _32

import (
	"container/heap"
)

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &maxHeap{}

	for i := range classes {
		*h = append(*h, classes[i])
	}

	heap.Init(h)

	for extraStudents > 0 {
		ratio := heap.Pop(h).([]int)

		ratio[0] += 1
		ratio[1] += 1

		heap.Push(h, ratio)

		extraStudents--
	}

	sum := float64(0)
	for _, ratio := range *h {
		sum += float64(ratio[0]) / float64(ratio[1])
	}

	return sum / float64(len(classes))
}

type maxHeap [][]int

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	if h[i][0] == h[i][1] {
		return false
	}

	if h[j][0] == h[j][1] {
		return true
	}

	left := (float64(h[i][0]+1) / float64(h[i][1]+1)) - (float64(h[i][0]) / float64(h[i][1]))
	right := (float64(h[j][0]+1) / float64(h[j][1]+1)) - (float64(h[j][0]) / float64(h[j][1]))

	return left > right
}
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *maxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}
