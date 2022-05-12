package _480_sliding_window_median

import (
	"container/heap"
)

// tptl. passed. complicated
func medianSlidingWindow(nums []int, k int) []float64 {
	s := StreamMedian{
		nums: nums,
		maxHeap: queue{
			less: func(i, j int) bool {
				return nums[i] > nums[j]
			},
			heapIDx: make(map[int]int),
		},
		minHeap: queue{
			less: func(i, j int) bool {
				return nums[i] < nums[j]
			},
			heapIDx: make(map[int]int),
		},
	}

	res := make([]float64, 0, len(nums)-k+1)
	for i := range nums {
		s.InsertNum(i)
		if i >= k {
			s.RemoveNum(i - k)
		}
		if i >= k-1 {
			res = append(res, s.FindMedian())
		}
	}

	return res
}

type StreamMedian struct {
	nums    []int
	maxHeap queue
	minHeap queue
}

func (s *StreamMedian) InsertNum(idx int) {
	if s.maxHeap.Len() == 0 || s.nums[s.maxHeap.idxes[0]] >= s.nums[idx] {
		heap.Push(&s.maxHeap, idx)
	} else {
		heap.Push(&s.minHeap, idx)
	}

	s.balance()
}

func (s *StreamMedian) FindMedian() float64 {
	if s.maxHeap.Len() == s.minHeap.Len() {
		num1, num2 := s.nums[s.maxHeap.idxes[0]], s.nums[s.minHeap.idxes[0]]
		return float64(num1+num2) / 2.0
	}

	return float64(s.nums[s.maxHeap.idxes[0]])
}

func (s *StreamMedian) RemoveNum(idx int) {
	if heapIDx, ok := s.maxHeap.heapIDx[idx]; ok {
		heap.Remove(&s.maxHeap, heapIDx)
	} else if heapIDx, ok = s.minHeap.heapIDx[heapIDx]; ok {
		heap.Remove(&s.minHeap, heapIDx)
	} else {
		panic("heapIDx not found")
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
	less    func(i, j int) bool
	idxes   []int
	heapIDx map[int]int
}

func (q queue) Len() int           { return len(q.idxes) }
func (q queue) Less(i, j int) bool { return q.less(q.idxes[i], q.idxes[j]) }
func (q queue) Swap(i, j int) {
	q.heapIDx[q.idxes[j]] = i
	q.heapIDx[q.idxes[i]] = j

	q.idxes[i], q.idxes[j] = q.idxes[j], q.idxes[i]
}

func (q *queue) Push(v interface{}) {
	num := v.(int)
	q.idxes = append(q.idxes, num)
	q.heapIDx[num] = len(q.idxes) - 1
}

func (q *queue) Pop() interface{} {
	last := len(q.idxes) - 1
	res := q.idxes[last]
	q.idxes = q.idxes[:last]

	delete(q.heapIDx, res)

	return res
}
