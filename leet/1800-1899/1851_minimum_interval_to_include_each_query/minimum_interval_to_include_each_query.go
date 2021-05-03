package _1851_minimum_interval_to_include_each_query

import (
	"container/heap"
	"sort"
)

// tptl. passed. hard
func minInterval(intervals [][]int, queries []int) []int {
	n := len(queries)
	res, sortedQueryIndex := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		sortedQueryIndex[i] = i
	}

	sort.Slice(sortedQueryIndex, func(i, j int) bool {
		return queries[sortedQueryIndex[i]] < queries[sortedQueryIndex[j]]
	})

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i, q := 0, &minHeap{}
	for j := 0; j < n; j++ {
		query := queries[sortedQueryIndex[j]]
		for i < len(intervals) && intervals[i][0] <= query {
			heap.Push(q, []int{intervals[i][1] - intervals[i][0] + 1, intervals[i][1]})
			i++
		}

		for q.Len() > 0 && (*q)[0][1] < query {
			heap.Pop(q)
		}

		if q.Len() == 0 {
			res[sortedQueryIndex[j]] = -1
		} else {
			res[sortedQueryIndex[j]] = (*q)[0][0]
		}
	}

	return res
}

type minHeap [][]int

func (h minHeap) Len() int              { return len(h) }
func (h minHeap) Less(i, j int) bool    { return h[i][0] < h[j][0] }
func (h minHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(val interface{}) { *h = append(*h, val.([]int)) }
func (h *minHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}
