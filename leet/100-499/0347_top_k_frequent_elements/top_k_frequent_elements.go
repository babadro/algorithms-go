package _347_top_k_frequent_elements

import (
	"container/heap"
	"sort"
)

// tptl. passed
func topKFrequent2(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	h := minHeap{freq: freq}
	for num, count := range freq {
		if h.Len() < k || freq[h.arr[0]] < count {
			heap.Push(&h, num)

			if h.Len() > k {
				heap.Pop(&h)
			}
		}
	}

	return h.arr
}

type minHeap struct {
	freq map[int]int
	arr  []int
}

func (h minHeap) Len() int            { return len(h.arr) }
func (h minHeap) Less(i, j int) bool  { return h.freq[h.arr[i]] < h.freq[h.arr[j]] }
func (h minHeap) Swap(i, j int)       { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }
func (h *minHeap) Push(v interface{}) { (*h).arr = append((*h).arr, v.(int)) }
func (h *minHeap) Pop() interface{} {
	last := len(h.arr) - 1
	res := h.arr[last]
	(*h).arr = (*h).arr[:last]

	return res
}

// bruteforce 93% 5%
func topKFrequent(nums []int, k int) []int {
	set := make(map[int]int, len(nums))

	for _, v := range nums {
		set[v]++
	}

	arr := make([][2]int, 0, len(set))

	for value, frequency := range set {
		arr = append(arr, [2]int{value, frequency})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})

	res := make([]int, k)

	for i := 0; i < k; i++ {
		res[i] = arr[i][0]
	}

	return res
}

// TODO 2 short solution (BucketSort)
// https://leetcode.com/problems/top-k-frequent-elements/discuss/371250/Go-MinHeap-and-Bucket-sort-solution
