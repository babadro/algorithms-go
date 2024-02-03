package _2316_count_Unreachable_pairs_of_node_in_undirected_graph

// #bnsrg medium passed
// todo 2 check faster solution (union find)
func countPairs(n int, edges [][]int) int64 {
	graph := make([][]int, n)

	for _, edge := range edges {
		src, dst := edge[0], edge[1]

		graph[src] = append(graph[src], dst)
		graph[dst] = append(graph[dst], src)
	}

	visited := make(map[int]bool)

	res := 0
	remainingNodes := n

	for vertex := range graph {
		prevCount := len(visited)

		if visited[vertex] {
			continue
		}

		dfs(graph, vertex, visited)

		componentSize := len(visited) - prevCount

		remainingNodes -= componentSize

		res += componentSize * remainingNodes
	}

	return int64(res)
}

func dfs(graph [][]int, vertex int, visited map[int]bool) {
	if visited[vertex] {
		return
	}

	visited[vertex] = true

	for _, adj := range graph[vertex] {
		dfs(graph, adj, visited)
	}
}
