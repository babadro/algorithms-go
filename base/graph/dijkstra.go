package graph

import (
	"container/heap"
	"math"
)

// https://www.programiz.com/dsa/dijkstra-algorithm
func Dijkstra(adjMatrix [][]int, vertexCount, start int) (distance []int) {
	cost := make([][]int, vertexCount)
	for i := range cost {
		cost[i] = make([]int, vertexCount)
	}

	distance, pred, visited := make([]int, vertexCount), make([]int, vertexCount), make([]bool, vertexCount)

	for i := 0; i < vertexCount; i++ {
		for j := 0; j < vertexCount; j++ {
			if adjMatrix[i][j] == 0 {
				cost[i][j] = math.MaxInt64
			} else {
				cost[i][j] = adjMatrix[i][j]
			}
		}
	}

	for i := 0; i < vertexCount; i++ {
		distance[i] = cost[start][i]
		pred[i] = start
	}

	distance[start] = 0
	visited[start] = true

	var nextV int
	for count := 1; count < vertexCount-1; count++ {
		minDistance := math.MaxInt64

		for i := 0; i < vertexCount; i++ {
			if distance[i] < minDistance && !visited[i] {
				minDistance = distance[i]
				nextV = i
			}
		}

		visited[nextV] = true
		for i := 0; i < vertexCount; i++ {
			if !visited[i] && adjMatrix[nextV][i] != 0 {
				if minDistance+cost[nextV][i] < distance[i] {
					distance[i] = minDistance + cost[nextV][i]
					pred[i] = nextV
				}
			}
		}
	}

	return distance
}

type MinHeap [][2]int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *MinHeap) Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

// todo 1 https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/discuss/1098349/Golang-Go-Dijkstra-with-min-heap-%2B-DFS-solution
func DijkstraWithHeap(adjList [][2]int, vertexCount, start int) (distance []int) {
	dist := make([]int, vertexCount+1)
	for i := range dist {
		dist[i] = math.MaxInt64
	}

	dist[start] = 0
	h := &MinHeap{}
	heap.Init(h)

	return nil
}
