package _31

import "github.com/babadro/algorithms-go/base/graph"

// todo 1 https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/discuss/1098310/C%2B%2B-BFS-%2B-DFS
func countRestrictedPaths(n int, edges [][]int) int {
	adjMatrix := make([][]int, n)
	for i := range adjMatrix {
		adjMatrix[i] = make([]int, n)
	}

	for _, edge := range edges {
		u, v, cost := edge[0]-1, edge[1]-1, edge[2]
		adjMatrix[u][v] = cost
		adjMatrix[v][u] = cost
	}

	distances := graph.Dijkstra(adjMatrix, n, n-1)

	dag := graph.New(n)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		if distances[u] == distances[v] {
			continue
		}

		if distances[u] > distances[v] {
			dag.AddEdge(u, v)
		} else {
			dag.AddEdge(v, u)
		}
	}

	paths := dag.AllPathsFromSourceToTarget(0, n-1)

	return len(paths)
}
