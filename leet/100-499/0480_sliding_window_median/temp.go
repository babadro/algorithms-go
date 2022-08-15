package _480_sliding_window_median

package _480_sliding_window_median

import "container/heap"

// tptl. passed. complicated
func medianSlidingWindow(nums []int, k int) []float64 {
	stream := streamMedian{}

	var res []float64
	for i := range nums {
		stream.insert(i)
		if i >= k {
			stream.remove(i - k)
		}

		if i >= k-1 {
			res = append(res, stream.findMedian())
		}
	}

	return res
}

type streamMedian struct {
	minHeap, maxHeap queue
	nums             []int
}

func (s *streamMedian) findMedian() float64 {
	return 0
}

func (s *streamMedian) insert(idx int) {
	num := s.nums[idx]
}

func (s *streamMedian) remove(idx int) {}

func (s *streamMedian) balance() {
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

func (q queue) Len() int {
	return len(q.numIDs)
}

func (q queue) Less(i, j int) bool {
	return q.less(q.numIDs[i], q.numIDs[j])
}

func (q queue) Swap(i, j int) {
	q.numIDxToHeapIDx[q.numIDs[i]] = j
	q.numIDxToHeapIDx[q.numIDs[j]] = i

	q.numIDs[i], q.numIDs[j] = q.numIDs[j], q.numIDs[j]
}

func (q *queue) Push(v interface{}) {
	idx := v.(int)
	q.numIDs = append(q.numIDs, idx)
	q.numIDxToHeapIDx[idx] = len(q.numIDs) - 1
}

func (q *queue) Pop() interface{} {
	last := q.Len() - 1

	numIDx := q.numIDs[last]
	q.numIDs = q.numIDs[:last]
	delete(q.numIDxToHeapIDx, numIDx)

	return numIDx
}

