package _2101_detonate_the_maximum_bombs

// bnsrg
func maximumDetonation(bombs [][]int) int {
	graph := make(map[int][]int, len(bombs))

	for j := range bombs {
		for i := range bombs {
			if i == j {
				continue
			}

			if distance(bombs[j], bombs[i]) <= bombs[j][2]*bombs[j][2] {
				graph[j] = append(graph[j], i)
			}
		}
	}

	res := 1
	for vertex := range graph {
		visited := make(map[int]bool)
		dfs(graph, vertex, visited)
		res = max(res, len(visited))
	}

	return res
}

func dfs(g map[int][]int, vertex int, visited map[int]bool) {
	if visited[vertex] {
		return
	}

	visited[vertex] = true

	for _, vertex := range g[vertex] {
		dfs(g, vertex, visited)
	}
}

func distance(b1, b2 []int) int {
	return (b1[0]-b2[0])*(b1[0]-b2[0]) + (b1[1]-b2[1])*(b1[1]-b2[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
