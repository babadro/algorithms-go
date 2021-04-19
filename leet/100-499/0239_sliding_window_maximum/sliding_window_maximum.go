package _239_sliding_window_maximum

import "container/heap"

type item struct {
	num     int
	numsIdx int
}

// passed. tptl hard todo 2 look for better solution
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n-k+1)

	h := &maxHeap{numsIdxToHeapIdx: make(map[int]int, k)}
	for i := 0; i < k; i++ {
		elem := item{
			num:     nums[i],
			numsIdx: i,
		}

		heap.Push(h, elem)
	}

	for i := range res {
		res[i] = h.arr[0].num

		next := i + k
		if next < n {
			idxToRemove := h.numsIdxToHeapIdx[i]
			heap.Remove(h, idxToRemove)

			heap.Push(h, item{nums[next], next})
		}

	}

	return res
}

type maxHeap struct {
	arr              []item
	numsIdxToHeapIdx map[int]int
}

func (h maxHeap) Len() int           { return len(h.arr) }
func (h maxHeap) Less(i, j int) bool { return h.arr[i].num > h.arr[j].num }

func (h maxHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
	h.numsIdxToHeapIdx[h.arr[i].numsIdx] = i
	h.numsIdxToHeapIdx[h.arr[j].numsIdx] = j
}

func (h *maxHeap) Push(x interface{}) {
	el := x.(item)
	h.arr = append(h.arr, el)
	h.numsIdxToHeapIdx[el.numsIdx] = len(h.arr) - 1
}

func (h *maxHeap) Pop() interface{} {
	n := len(h.arr)
	res := h.arr[n-1]
	h.arr = h.arr[:n-1]

	delete(h.numsIdxToHeapIdx, res.numsIdx)

	return res
}
