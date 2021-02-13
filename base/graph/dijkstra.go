package graph

import "math"

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
