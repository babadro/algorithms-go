// dyx
package graph

import (
	"container/list"
)

type Graph struct {
	adj [][]int
}

func New(V int) *Graph {
	return &Graph{make([][]int, V)}
}

func Constructor(g [][]int) *Graph {
	return &Graph{g}
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
	q := list.New()

	visited[v] = true
	q.PushBack(v)

	for q.Front() != nil {
		v = q.Front().Value.(int)
		q.Remove(q.Front())

		f(v)

		for _, adjV := range g.adj[v] {
			if !visited[adjV] {
				visited[adjV] = true
				q.PushBack(adjV)
			}
		}
	}
}

// dyx
// todo base add unit tests
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

// graph must be direct and acyclic (DAG)
// todo 1 add unit tests where src and target would be not equal 0 and n-1.
func (g *Graph) AllPathsFromSourceToTarget(source, target int, f func(path []int)) {
	path := make([]int, 0)

	g.allPathsHelper(source, target, path, f)
}

func (g *Graph) allPathsHelper(V, target int, path []int, f func(path []int)) {
	path = append(path, V)

	if V == target {
		f(path)
		return
	}

	for _, v := range g.adj[V] {
		g.allPathsHelper(v, target, path, f)
	}
}
