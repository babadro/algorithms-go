package adjgraph

import (
	"fmt"
	"github.com/babadro/algorithms-go/03_StacksAndQueues/03_Queue"
	"strings"
)

type Graph struct {
	V   int
	adj [][]int
}

func New(V int) *Graph {
	return &Graph{V, make([][]int, V)}
}

// Direction graph
func (g *Graph) AddEdge(src, dst int) {
	g.adj[src] = append(g.adj[src], dst)
	// g.adj[dst] = append(g.adj[dst], src) non direction graph
}

func (g *Graph) BFS(v int) {
	visited := make([]bool, g.V)

	q := queue.New(g.V)

	visited[v] = true
	q.Enqueue(v)

	for q.Len() != 0 {
		v, _ = q.Dequeue()
		fmt.Print(v, " ")

		for _, n := range g.adj[v] {
			if !visited[n] {
				visited[n] = true
				q.Enqueue(n)
			}
		}
	}
}

func (g *Graph) DFS(i int) {
	visited := make([]bool, g.V)
	g.dfsUtil(i, visited)
}

func (g *Graph) dfsUtil(v int, visited []bool) {
	visited[v] = true
	fmt.Print(v, " ")

	vList := g.adj[v]
	for _, n := range vList {
		if !visited[n] {
			g.dfsUtil(n, visited)
		}
	}
}

func (g *Graph) String() string {
	var str strings.Builder
	var v int
	for v = 0; v < g.V; v++ {
		pCrawl := g.adj[v]
		str.WriteString(fmt.Sprintf("\n Ajacency list of n %d\n head ", v))
		for _, n := range pCrawl {
			str.WriteString(fmt.Sprintf("-> %d", n))
		}
		str.WriteString("\n")
	}
	return str.String()
}
