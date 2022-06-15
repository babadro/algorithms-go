package _1_topological_sort

// leetcode 210
func topologicalSort(vertices int, edges [][]int) []int {
	inDegree := make(map[int]int, vertices)
	for _, edge := range edges {
		inDegree[edge[1]]++
	}

	// build the graph
	g := make(map[int][]int, vertices)
	for i := 0; i < vertices; i++ {
		g[i] = nil
	}
	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		g[src] = append(g[src], dst)
	}

	var sources []int
	for vertex := range g {
		if _, ok := inDegree[vertex]; !ok {
			sources = append(sources, vertex)
		}
	}

	for i := 0; i < len(sources); i++ {
		src := sources[i]

		for _, dst := range g[src] {
			inDegree[dst]--
			if inDegree[dst] == 0 {
				sources = append(sources, dst)
			}
		}
	}

	return sources
}
