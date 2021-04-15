package _1825_finding_mk_average

import (
	"container/heap"
)

// MKAverageHeap doesn't work (TLE), but, probably, correct
type MKAverageHeap struct {
	nums   []int
	m      int
	k      int
	maxH   *maxHeap
	minH   *minHeap
	minTmp []item
	maxTmp []item
	lenRes int
}

func Constructor(m int, k int) MKAverageHeap {
	maxH := &maxHeap{numsIdxToHeapIdx: make(map[int]int, m)}
	minH := &minHeap{numsIdxToHeapIdx: make(map[int]int, m)}

	return MKAverageHeap{
		m:      m,
		k:      k,
		maxH:   maxH,
		minH:   minH,
		lenRes: m - 2*k,
	}
}

func (mk *MKAverageHeap) AddElement(num int) {
	mk.nums = append(mk.nums, num)

	elem := item{
		num:     num,
		numsIdx: len(mk.nums) - 1,
	}

	if len(mk.nums) > mk.m {
		numsIdx := len(mk.nums) - 1 - mk.m

		minHIDx := mk.minH.numsIdxToHeapIdx[numsIdx]
		heap.Remove(mk.minH, minHIDx)

		maxHIDx := mk.maxH.numsIdxToHeapIdx[numsIdx]
		heap.Remove(mk.maxH, maxHIDx)
	}

	heap.Push(mk.minH, elem)
	heap.Push(mk.maxH, elem)
}

func (mk *MKAverageHeap) CalculateMKAverage() int {
	if len(mk.nums) < mk.m {
		return -1
	}

	sum := 0
	for i := len(mk.nums) - mk.m; i < len(mk.nums); i++ {
		sum += mk.nums[i]
	}

	mk.minTmp = mk.minTmp[:0]
	mk.maxTmp = mk.maxTmp[:0]
	for i := 0; i < mk.k; i++ {
		min := heap.Pop(mk.minH).(item)
		max := heap.Pop(mk.maxH).(item)

		sum -= min.num
		sum -= max.num

		mk.minTmp = append(mk.minTmp, min)
		mk.maxTmp = append(mk.maxTmp, max)
	}

	for i := 0; i < mk.k; i++ {
		heap.Push(mk.minH, mk.minTmp[i])
		heap.Push(mk.maxH, mk.maxTmp[i])
	}

	return sum / mk.lenRes
}

type item struct {
	num     int
	numsIdx int
}

type minHeap struct {
	arr              []item
	numsIdxToHeapIdx map[int]int
}

func (h minHeap) Len() int           { return len(h.arr) }
func (h minHeap) Less(i, j int) bool { return h.arr[i].num < h.arr[j].num }

func (h minHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
	h.numsIdxToHeapIdx[h.arr[i].numsIdx] = i
	h.numsIdxToHeapIdx[h.arr[j].numsIdx] = j
}

func (h *minHeap) Push(x interface{}) {
	el := x.(item)
	h.arr = append(h.arr, el)
	h.numsIdxToHeapIdx[el.numsIdx] = len(h.arr) - 1
}

func (h *minHeap) Pop() interface{} {
	n := len(h.arr)
	res := h.arr[n-1]
	h.arr = h.arr[:n-1]

	delete(h.numsIdxToHeapIdx, res.numsIdx)

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
