package _3_calculcate_the_number_of_nodes_in_a_given_graph_level

type graph [][]int

// tptl. edu #graph
func Calculate(g [][]int, level int) int {
	newGraph := graph(g)

	return newGraph.NumberOfNodes(level)
}

func (g *graph) NumberOfNodes(level int) int {
	visited := make(map[int]bool)
	var count int
	g.dfs(visited, 0, 1, level, &count)

	return count
}

func (g *graph) dfs(visited map[int]bool, v, currLevel, targetLevel int, count *int) {
	if currLevel == targetLevel {
		*count += 1
		return
	}

	for _, adj := range (*g)[v] {
		if !visited[adj] {
			visited[adj] = true
			g.dfs(visited, adj, currLevel+1, targetLevel, count)
		}
	}
}
