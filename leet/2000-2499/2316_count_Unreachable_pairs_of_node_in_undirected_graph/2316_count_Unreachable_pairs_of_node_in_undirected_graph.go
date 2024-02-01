package _2316_count_Unreachable_pairs_of_node_in_undirected_graph

// #bnsrg medium todo 1 - tle
func countPairs(n int, edges [][]int) int64 {
	graph := make(map[int][]int)

	for i := 0; i < n; i++ {
		graph[i] = nil
	}

	for _, edge := range edges {
		src, dst := edge[0], edge[1]

		graph[src] = append(graph[src], dst)
		graph[dst] = append(graph[dst], src)
	}

	var connectedGraphLenghts []int

	visited := make(map[int]bool)

	for vertex := range graph {
		prevCount := len(visited)

		if visited[vertex] {
			continue
		}

		dfs(graph, vertex, visited)

		connectedGraphLen := len(visited) - prevCount

		connectedGraphLenghts = append(connectedGraphLenghts, connectedGraphLen)
	}

	res := 0
	for i := 0; i < len(connectedGraphLenghts)-1; i++ {
		for j := i + 1; j < len(connectedGraphLenghts); j++ {
			res += connectedGraphLenghts[i] * connectedGraphLenghts[j]
		}
	}

	return int64(res)
}

func dfs(graph map[int][]int, vertex int, visited map[int]bool) {
	if visited[vertex] {
		return
	}

	visited[vertex] = true

	for _, adj := range graph[vertex] {
		dfs(graph, adj, visited)
	}
}
