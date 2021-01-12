// dyx
package graph

import "container/list"

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

// dyx
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

// dyx
// todo 1 add unit tests
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

// dyx
func (g *Graph) TopologicalSort(f func(v int)) {
	stack := make([]int, 0)
	visited := make(map[int]bool)

	for i := 0; i < g.V(); i++ {
		if !visited[i] {
			g.topologicalSortHelper(i, visited, &stack)
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		f(stack[i])
	}
}

func (g *Graph) topologicalSortHelper(v int, visited map[int]bool, stack *[]int) {
	visited[v] = true
	for _, adjV := range g.adj[v] {
		if !visited[adjV] {
			g.topologicalSortHelper(adjV, visited, stack)
		}
	}

	*stack = append(*stack, v)
}
