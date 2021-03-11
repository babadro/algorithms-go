package _1786_number_of_restricted_paths_from_first_to_last_node

import "github.com/babadro/algorithms-go/base/graph"

// passed. dyx.
func countRestrictedPaths3(n int, edges [][]int) int {
	adjList := graph.EdgesSliceToAdjList(n, edges, true)
	distances := graph.DijkstraWithHeap(adjList, n, n-1)

	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	return dfs2(0, memo, adjList, distances, n-1)
}

func dfs2(u int, memo []int, graph [][][2]int, dist []int, start int) int {
	if memo[u] >= 0 {
		return memo[u]
	}

	if u == start {
		return 1
	}

	memo[u] = 0
	for _, v := range graph[u] {
		vVertex := v[1]
		if dist[vVertex] < dist[u] {
			memo[u] = (memo[u] + dfs2(vVertex, memo, graph, dist, start)) % 1_000_000_007
		}
	}

	return memo[u]
}
