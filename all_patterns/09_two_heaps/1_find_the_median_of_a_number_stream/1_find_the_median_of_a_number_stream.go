package _1_find_the_median_of_a_number_stream

import (
	"container/heap"

	"github.com/babadro/algorithms-go/utils"
)

// leetcode 295
type StreamMedian struct {
	maxHeap utils.IntMaxHeap
	minHeap utils.IntMinHeap
}

func (s *StreamMedian) InsertNum(num int) {
	if len(s.maxHeap) == 0 || s.maxHeap[0] >= num {
		heap.Push(&s.maxHeap, num)
	} else {
		heap.Push(&s.minHeap, num)
	}

	// either both the heaps will have equal number of elements or max-heap will have one
	// more element than the min-heap
	if len(s.maxHeap) > len(s.minHeap)+1 {
		heap.Push(&s.minHeap, heap.Pop(&s.maxHeap))
	} else if len(s.maxHeap) < len(s.minHeap) {
		heap.Push(&s.maxHeap, heap.Pop(&s.minHeap))
	}
}

func (s *StreamMedian) FindMedian() float64 {
	if len(s.maxHeap) == len(s.minHeap) {
		return float64(s.maxHeap[0]+s.minHeap[0]) / 2.0
	}

	return float64(s.maxHeap[0])
}
