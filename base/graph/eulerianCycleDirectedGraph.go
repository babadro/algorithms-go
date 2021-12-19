package graph

// https://www.techiedelight.com/eulerian-cycle-directed-graph/ has eulerian cycle todo base
func (g *EulerianGraph) HasEulerianCycle() bool {
	// check if every vertex has the same in-degree and out-degree
	for i := range g.adj {
		if len(g.adj[i]) != g.in[i] {
			return false
		}
	}

	// check if all vertices with a non-zero degree belong to a single
	// strongly connected component
	return IsStrongConnected(g.Graph)
}

// there is more effective algorithm here todo base
// https://www.techiedelight.com/check-graph-strongly-connected-one-dfs-traversal
func IsStrongConnected(g Graph) bool {
	var startVertex int

	var visited map[int]bool
	for i := range g.adj {
		if len(g.adj[i]) != 0 {
			startVertex = i
			visited = g.DFS(startVertex, func(v int) {})

			break
		}
	}

	if visited == nil {
		panic("graph must contain at least one connected component")
	}

	if !isVisited(g, visited) {
		return false
	}

	transposedGraph := BuildTransposed(g)
	visited = transposedGraph.DFS(startVertex, func(v int) {})
	if visited == nil {
		panic("graph must contain at least one connected component")
	}

	return isVisited(transposedGraph, visited)
}

func isVisited(g Graph, visited map[int]bool) bool {
	for i := range g.adj {
		if len(g.adj[i]) != 0 && !visited[i] {
			return false
		}
	}

	return true
}

func BuildTransposed(g Graph) Graph {
	transposedGraph := New(g.V())
	for src := range g.adj {
		for dst := range g.adj[src] {
			transposedGraph.AddEdge(dst, src)
		}
	}

	return *transposedGraph
}
