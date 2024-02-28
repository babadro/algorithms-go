package _834_sum_of_distances_in_tree

// #bnsrg hard passed
func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make([][]int, n)

	for _, edge := range edges {
		src, dst := edge[0], edge[1]

		graph[src] = append(graph[src], dst)
		graph[dst] = append(graph[dst], src)
	}

	ans, count := make([]int, n), make([]int, n)

	for i := range count {
		count[i] = 1
	}

	dfs1(graph, ans, count, 0, -1)
	dfs2(graph, ans, count, 0, -1, n)

	return ans
}

func dfs1(graph [][]int, ans, count []int, node, parent int) {
	for _, child := range graph[node] {
		if child != parent {
			dfs1(graph, ans, count, child, node)

			count[node] += count[child]
			ans[node] += ans[child] + count[child]
		}
	}
}

func dfs2(graph [][]int, ans, count []int, node, parent, n int) {
	for _, child := range graph[node] {
		if child != parent {
			ans[child] = n - 2*count[child] + ans[node]

			dfs2(graph, ans, count, child, node, n)
		}
	}
}
