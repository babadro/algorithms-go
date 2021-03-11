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
// todo 1 - find more effective solution. add unit tests where src and target would be not equal 0 and n-1.
func (g *Graph) AllPathsFromSourceToTarget(source, target int) [][]int {
	return g.allPathsHelper(source, target, make([][][]int, len(g.adj)))
}

func (g *Graph) allPathsHelper(source, target int, cache [][][]int) [][]int {
	if source == target {
		return [][]int{{source}}
	}

	if cache[source] != nil {
		return cache[source]
	}

	var res [][]int

	for _, v := range g.adj[source] {
		paths := g.allPathsHelper(v, target, cache)

		for _, childPath := range paths {
			path := make([]int, 0, len(childPath)+1)
			path = append(path, source)
			path = append(path, childPath...)

			res = append(res, path)
		}
	}

	cache[source] = res

	return res
}

// todo 1 doesn't work. check it
func (g *Graph) AllPathsFromSourceToTarget2(source, target int) [][]int {
	sortedVertexes := make([]int, 0, g.V())
	f := func(v int) {
		sortedVertexes = append(sortedVertexes, v)
	}

	g.TopologicalSort(f)

	d := make([][][]int, 0)

	n := g.V()
	for _, l := range sortedVertexes {
		if l == n {
			d[l] = [][]int{{n - 1}}
		} else {
			tmp := make([][]int, 0)
			for _, v := range g.adj[l] {
				for _, route := range d[v] {
					item := append([]int{l}, route...)
					tmp = append(tmp, item)
				}
				d[l] = tmp
			}
		}
	}

	return d[0]
}

/*
func (g *Graph) AllPathsFromSourceToTarget(source, target int) [][]int {
	sortedVertexes := make([]int, 0, g.V())
	f := func(v int) {
		sortedVertexes = append(sortedVertexes, v)
	}

	g.TopologicalSort(f)

	dp := make([]int, g.V())

	dp[target] = 1

	for i := len(sortedVertexes)-1; i >= 0; i-- {
		vertex := sortedVertexes[i]
		for j := 0; j < len(g.adj[vertex]); i-- {
			dp[vertex] += dp[g.adj[vertex][j]]
		}
	}

	return dp
}
*/
