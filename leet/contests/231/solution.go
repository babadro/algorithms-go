package _31

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

// todo 1 fails on 422 input
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
