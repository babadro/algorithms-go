package _2203_minimum_weighted_subgraph_with_the_required_paths

import (
	"container/heap"
	"math"
)

// https://leetcode.com/problems/minimum-weighted-subgraph-with-the-required-paths/discuss/1844095/Three-Dijkstras
// todo 1 doesn't work - need to recheck. #hard
func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	maxVal, res := 10000000000, math.MaxInt64
	al, ral := make([][][2]int, n), make([][][2]int, n)
	dd, s1d, s2d := makeSlice(n, maxVal), makeSlice(n, maxVal), makeSlice(n, maxVal)
	dd[dest], s1d[src1], s2d[src2] = 0, 0, 0

	for _, e := range edges {
		from, to, weight := e[0], e[1], e[2]

		al[from] = append(al[from], [2]int{to, weight})
		ral[to] = append(ral[to], [2]int{from, weight})
	}

	bfs(dest, ral, dd)
	bfs(src1, al, s1d)
	bfs(src2, al, s2d)

	if dd[src1] == maxVal || dd[src2] == maxVal {
		return -1
	}

	for i := 0; i < n; i++ {
		res = min(res, dd[i]+s1d[i]+s2d[i])
	}

	return int64(res)
}

func bfs(st int, al [][][2]int, visited []int) {
	pq := minHeap{}
	heap.Push(&pq, [2]int{st, 0})

	for len(pq) != 0 {
		top := heap.Pop(&pq).([2]int)
		dist, i := top[0], top[1]

		if visited[i] != dist {
			continue
		}

		for _, edge := range al[i] {
			j, w := edge[0], edge[1]
			if visited[j] > dist+w {
				visited[j] = dist + w
				heap.Push(&pq, [2]int{visited[j], j})
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
