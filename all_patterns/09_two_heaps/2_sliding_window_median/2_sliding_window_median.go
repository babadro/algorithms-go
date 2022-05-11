package __sliding_window_median

import "container/heap"

// leetcode 480
func medianSlidingWindow(nums []int, k int) []float64 {
	s := StreamMedian{
		maxHeap: queue{
			less: func(i, j int) bool {
				return nums[i] > nums[j]
			},
			numToIdx: make(map[int]int),
		},
		minHeap: queue{
			less: func(i, j int) bool {
				return nums[i] < nums[j]
			},
			numToIdx: make(map[int]int),
		},
	}

}

type StreamMedian struct {
	maxHeap queue
	minHeap queue
}

func (s *StreamMedian) InsertNum(num int) {
	if s.maxHeap.Len() == 0 || s.maxHeap.nums[0] >= num {
		heap.Push(&s.maxHeap, num)
	} else {
		heap.Push(&s.minHeap, num)
	}

	s.balance()
}

func (s *StreamMedian) FindMedian() float64 {
	if s.maxHeap.Len() == s.minHeap.Len() {
		return float64(s.maxHeap.nums[0]+s.minHeap.nums[0]) / 2.0
	}

	return float64(s.maxHeap.nums[0])
}

func (s *StreamMedian) RemoveNum(num int) {
	if idx, ok := s.maxHeap.numToIdx[num]; ok {
		heap.Remove(&s.maxHeap, idx)
	} else if idx, ok = s.minHeap.numToIdx[num]; ok {
		heap.Remove(&s.minHeap, idx)
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
	less     func(i, j int) bool
	nums     []int
	numToIdx map[int]int
}

func (q queue) Len() int           { return len(q.nums) }
func (q queue) Less(i, j int) bool { return q.less(q.nums[i], q.nums[j]) }
func (q queue) Swap(i, j int) {
	q.numToIdx[q.nums[j]] = i
	q.numToIdx[q.nums[i]] = j

	q.nums[i], q.nums[j] = q.nums[j], q.nums[i]
}

func (q *queue) Push(v interface{}) {
	num := v.(int)
	q.nums = append(q.nums, num)
	q.numToIdx[num] = len(q.nums) - 1
}

func (q *queue) Pop() interface{} {
	last := len(q.nums) - 1
	res := q.nums[last]
	q.nums = q.nums[:last]

	delete(q.numToIdx, res)

	return res
}
