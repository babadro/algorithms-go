package graph

// directed graph
// https://www.techiedelight.com/eulerian-path-directed-graph/	has eulerian path todo base

type EulerianPahGraph struct {
	// to store in-degree of each vertex in the graph
	in []int

	Graph
}

func NewEulerianPathGraph(V int) *EulerianPahGraph {
	return &EulerianPahGraph{
		in:    make([]int, V),
		Graph: *New(V),
	}
}

func (g *EulerianPahGraph) AddEdge(u, v int) {
	g.AddEdge(u, v)
	g.in[v]++
}

func getUndirectedGraph(g Graph) Graph {
	undirectedGraph := New(g.V())
	uniqBackDirectionEdges := make(map[[2]int]bool)
	for i := 0; i < g.V(); i++ {
		for _, v := range g.adj[i] {
			undirectedGraph.AddEdge(i, v)

		}
	}

	return adjMatrix
}

func IsConnected(g Graph) bool {
	var visited map[int]bool

	// start DFS from the first vertex with a non-zero degree
	for i := 0; i < g.V(); i++ {
		if len(g.adj[i]) != 0 {
			visited = g.DFS(i, func(v int) {})

			break
		}
	}

	if visited == nil {
		panic("graph must contain at least one connected component")
	}

	// if a single DFS call couldn't visit all vertices with a non-zero degree,
	// the graph contains more than one connected component
	for i := 0; i < g.V(); i++ {
		if !visited[i] && len(g.adj[i]) != 0 {
			return false
		}
	}

	return true
}

func HasEulerPath(g EulerianPahGraph) bool {
	/*
	   The following loop checks the following conditions to determine if an
	   Eulerian path can exist or not:
	       a. At most one vertex in the graph has `out-degree = 1 + in-degree`.
	       b. At most one vertex in the graph has `in-degree = 1 + out-degree`.
	       c. Rest all vertices have `in-degree == out-degree`.

	   If either of the above condition fails, the Euler path can't exist.
	*/

	x, y := false, false
	for i := 0; i < g.V(); i++ {
		outDegree := len(g.adj[i])
		inDegree := g.in[i]

		if outDegree != inDegree {
			if !x && outDegree-inDegree == 1 {
				x = true
			} else if !y && inDegree-outDegree == 1 {
				y = true
			} else {
				return false
			}
		}

	}
}

// https://www.techiedelight.com/eulerian-cycle-directed-graph/ has eulerian cycle todo base

// undirected graph
// https://www.techiedelight.com/eulerian-path-undirected-graph/ has eulerian path or cycle todo base
