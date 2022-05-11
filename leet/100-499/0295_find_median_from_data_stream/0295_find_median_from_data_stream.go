package _295_find_median_from_data_stream

import (
	"container/heap"
)

// tptl passed
type MedianFinder struct {
	maxHeap IntMaxHeap
	minHeap IntMinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (s *MedianFinder) AddNum(num int) {
	if len(s.maxHeap) == 0 || s.maxHeap[0] >= num {
		heap.Push(&s.maxHeap, num)
	} else {
		heap.Push(&s.minHeap, num)
	}

	if len(s.maxHeap) > len(s.minHeap)+1 {
		heap.Push(&s.minHeap, heap.Pop(&s.maxHeap))
	} else if len(s.maxHeap) < len(s.minHeap) {
		heap.Push(&s.maxHeap, heap.Pop(&s.minHeap))
	}
}

func (s *MedianFinder) FindMedian() float64 {
	if len(s.maxHeap) == len(s.minHeap) {
		return float64(s.maxHeap[0]+s.minHeap[0]) / 2.0
	}

	return float64(s.maxHeap[0])
}

type IntMinHeap []int

func (h IntMinHeap) Len() int            { return len(h) }
func (h IntMinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *IntMinHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}

type IntMaxHeap []int

func (h IntMaxHeap) Len() int            { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntMaxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *IntMaxHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}
