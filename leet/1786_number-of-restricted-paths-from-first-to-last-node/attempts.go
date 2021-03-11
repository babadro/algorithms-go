package _1786_number_of_restricted_paths_from_first_to_last_node

import "github.com/babadro/algorithms-go/base/graph"

// doesn't work. Time limit exceeded.
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

// todo 1 fails on 422 input
// not mine: https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/discuss/1098310/C%2B%2B-BFS-%2B-DFS
func countRestrictedPaths2(n int, edges [][]int) int {
	al := make([][][2]int, n+1)
	dist := make([]int, n+1)

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}

	q := []int{n}

	for _, edge := range edges {
		u, v, cost := edge[0], edge[1], edge[2]

		al[u] = append(al[u], [2]int{v, cost})
		al[v] = append(al[v], [2]int{u, cost})
	}

	for len(q) > 0 {
		var q1 []int
		for _, i := range q {
			for _, pair := range al[i] {
				j, w := pair[0], pair[1]
				if j != n && (dist[j] == 0 || dist[j] > dist[i]+w) {
					dist[j] = dist[i] + w
					q1 = append(q1, j)
				}
			}
		}

		q, q1 = q1, q
	}

	return dfs(al, dist, dp, n)
}

func dfs(al [][][2]int, dist []int, dp []int, i int) int {
	if i == 1 {
		return 1
	}

	if dp[i] == -1 {
		dp[i] = 0
		for _, pair := range al[i] {
			j := pair[0]
			if dist[i] < dist[j] {
				dp[i] = (dp[i] + dfs(al, dist, dp, j)) % 1_000_007
			}
		}
	}

	return dp[i]
}
