package _3_calculcate_the_number_of_nodes_in_a_given_graph_level

import "container/list"

type graph [][]int

// tptl. edu #graph
// recursive
func CalculateDFS(g [][]int, level int) int {
	newGraph := graph(g)

	return newGraph.NumberOfNodes(level)
}

// iterative
func CalculateBFS(g [][]int, level int) int {
	newGraph := graph(g)

	return newGraph.NumberOfNodesBFS(level)
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

func (g *graph) NumberOfNodesBFS(level int) int {
	visited := make(map[int]int, len(*g))
	q := list.New()
	q.PushBack(0)

	curLevel := 1
	for q.Len() > 0 {
		v := q.Front().Value.(int)
		q.Remove(q.Front())

		for _, adj := range (*g)[v] {
			if visited[adj] == 0 {
				visited[adj] = curLevel + 1
				q.PushBack(adj)
			}
		}

		curLevel++
	}

	res := 0
	for _, v := range visited {
		if v == level {
			res++
		}
	}

	return res
}
