package _1857_largest_color_value_in_a_directed_graph

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

// todo 1 move to base graph library
func (g *Graph) isCyclicUtil(i int, visited, recStack []bool) bool {
	if recStack[i] {
		return true
	}

	if visited[i] {
		return false
	}

	visited[i] = true

	recStack[i] = true

	children := g.adj[i]

	for _, v := range children {
		if g.isCyclicUtil(v, visited, recStack) {
			return true
		}
	}

	recStack[i] = false

	return false
}

func (g *Graph) IsCyclic() bool {
	visited := make([]bool, g.V())
	recStack := make([]bool, g.V())

	for i := 0; i < g.V(); i++ {
		if g.isCyclicUtil(i, visited, recStack) {
			return true
		}
	}

	return false
}
