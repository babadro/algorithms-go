package _480_sliding_window_median

import (
	"container/heap"
)

func medianSlidingWindow(nums []int, k int) []float64 {
	s := StreamMedian{
		nums: nums,
		maxHeap: queue{
			less: func(i, j int) bool {
				return nums[i] > nums[j]
			},
			numIDxToHeapIDx: make(map[int]int),
		},
		minHeap: queue{
			less: func(i, j int) bool {
				return nums[i] < nums[j]
			},
			numIDxToHeapIDx: make(map[int]int),
		},
	}

	res := make([]float64, 0, len(nums)-k+1)
	for i := range nums {
		s.insert(i)
		if i >= k {
			s.remove(i - k)
		}
		if i >= k-1 {
			res = append(res, s.findMedian())
		}
	}

	return res
}

type StreamMedian struct {
	nums    []int
	maxHeap queue
	minHeap queue
}

func (s *StreamMedian) insert(idx int) {
	if s.maxHeap.Len() == 0 || s.nums[s.maxHeap.numIDs[0]] >= s.nums[idx] {
		heap.Push(&s.maxHeap, idx)
	} else {
		heap.Push(&s.minHeap, idx)
	}

	s.balance()
}

func (s *StreamMedian) findMedian() float64 {
	if s.maxHeap.Len() == s.minHeap.Len() {
		num1, num2 := s.nums[s.maxHeap.numIDs[0]], s.nums[s.minHeap.numIDs[0]]
		return float64(num1+num2) / 2.0
	}

	return float64(s.nums[s.maxHeap.numIDs[0]])
}

func (s *StreamMedian) remove(idx int) {
	if idxOfIdx, ok := s.maxHeap.numIDxToHeapIDx[idx]; ok {
		heap.Remove(&s.maxHeap, idxOfIdx)
	} else if idxOfIdx, ok = s.minHeap.numIDxToHeapIDx[idx]; ok {
		heap.Remove(&s.minHeap, idxOfIdx)
	} else {
		panic("idx not found")
	}

	s.balance()
}

func (s *StreamMedian) balance() {
	if s.maxHeap.Len() > s.minHeap.Len()+1 {
		heap.Push(&s.minHeap, heap.Pop(&s.maxHeap))
	} else if s.maxHeap.Len() < s.minHeap.Len() {
		heap.Push(&s.maxHeap, heap.Pop(&s.minHeap))
	}
}

type queue struct {
	less            func(i, j int) bool
	numIDs          []int
	numIDxToHeapIDx map[int]int
}

func (q queue) Len() int           { return len(q.numIDs) }
func (q queue) Less(i, j int) bool { return q.less(q.numIDs[i], q.numIDs[j]) }
func (q queue) Swap(i, j int) {
	q.numIDxToHeapIDx[q.numIDs[j]] = i
	q.numIDxToHeapIDx[q.numIDs[i]] = j

	q.numIDs[i], q.numIDs[j] = q.numIDs[j], q.numIDs[i]
}

func (q *queue) Push(v interface{}) {
	num := v.(int)
	q.numIDs = append(q.numIDs, num)
	q.numIDxToHeapIDx[num] = len(q.numIDs) - 1
}

func (q *queue) Pop() interface{} {
	last := len(q.numIDs) - 1
	res := q.numIDs[last]
	q.numIDs = q.numIDs[:last]

	delete(q.numIDxToHeapIDx, res)

	return res
}
