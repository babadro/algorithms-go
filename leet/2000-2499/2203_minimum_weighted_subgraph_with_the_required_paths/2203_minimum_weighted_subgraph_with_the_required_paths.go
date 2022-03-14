package _2203_minimum_weighted_subgraph_with_the_required_paths

import (
	"container/heap"
	"math"
)

// https://leetcode.com/problems/minimum-weighted-subgraph-with-the-required-paths/discuss/1844095/Three-Dijkstras
// dyx. todo 2 understand. #hard
// there is a more simple version without priority queue: https://leetcode.com/problems/minimum-weighted-subgraph-with-the-required-paths/discuss/1846504/Go-GoLang-or-Time%3A-O(n-log-n-*-3)-355-ms-100-or-Space%3A-O(n-*-5)-31.1-MB-100
func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	maxVal, res := 10000000000, math.MaxInt64
	graph, reversedGraph := make([][][2]int, n), make([][][2]int, n)
	commonDist, src1Dist, src2Dist := makeSlice(n, maxVal), makeSlice(n, maxVal), makeSlice(n, maxVal)
	commonDist[dest], src1Dist[src1], src2Dist[src2] = 0, 0, 0

	for _, e := range edges {
		from, to, weight := e[0], e[1], e[2]

		graph[from] = append(graph[from], [2]int{to, weight})
		reversedGraph[to] = append(reversedGraph[to], [2]int{from, weight})
	}

	bfs(dest, reversedGraph, commonDist)
	bfs(src1, graph, src1Dist)
	bfs(src2, graph, src2Dist)

	if commonDist[src1] == maxVal || commonDist[src2] == maxVal {
		return -1
	}

	for i := 0; i < n; i++ {
		res = min(res, commonDist[i]+src1Dist[i]+src2Dist[i])
	}

	return int64(res)
}

func bfs(st int, al [][][2]int, visited []int) {
	priorityQ := minHeap{}
	heap.Push(&priorityQ, [2]int{0, st})

	for len(priorityQ) != 0 {
		top := heap.Pop(&priorityQ).([2]int)
		dist, i := top[0], top[1]

		if visited[i] != dist {
			continue
		}

		for _, edge := range al[i] {
			j, w := edge[0], edge[1]
			if visited[j] > dist+w {
				visited[j] = dist + w
				heap.Push(&priorityQ, [2]int{visited[j], j})
			}
		}

	}
}

type minHeap [][2]int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *minHeap) Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func makeSlice(length, val int) []int {
	res := make([]int, length)
	for i := range res {
		res[i] = val
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
