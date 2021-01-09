package graph

import "container/list"

// todo 1 make unit tests
type Graph struct {
	adj [][]int
}

func New(V int) *Graph {
	return &Graph{make([][]int, V)}
}

func (g *Graph) AddEdge(src, dst int) {
	g.adj[src] = append(g.adj[src], dst)
	//g.adj[dst] = append(g.adj[dst], src) // non directional graph
}

func (g *Graph) V() int {
	return len(g.adj)
}

func (g *Graph) BFS(v int, f func(v int)) {
	visited := make(map[int]bool, g.V())
	visited[v] = true

	q := list.New()
	q.PushBack(v)

	for q.Front() != nil {
		v = q.Front().Value.(int)
		q.Remove(q.Front())

		f(v)

		for _, adjV := range g.adj[v] {
			if !visited[adjV] {
				q.PushBack(adjV)
				visited[adjV] = true
			}
		}
	}
}

func (g *Graph) DFS(v int, f func(v int)) {
	visited := make(map[int]bool, g.V())
	g.dfsHelper(v, visited, f)
}

func (g *Graph) dfsHelper(v int, visited map[int]bool, f func(v int)) {
	visited[v] = true

	f(v)

	for _, adjV := range g.adj[v] {
		if !visited[adjV] {
			g.dfsHelper(adjV, visited, f)
		}
	}
}
