package adjgraph

import (
	"fmt"
	"github.com/babadro/algorithms-go/03_StacksAndQueues/03_Queue"
	"strings"
)

type AdjListNode struct {
	Vertex int
	Next   *AdjListNode
}

func NewAdjListNode(dest int) *AdjListNode {
	return &AdjListNode{dest, nil}
}

type Graph struct {
	V     int
	Array []*AdjListNode
}

func New(V int) *Graph {
	return &Graph{V, make([]*AdjListNode, V)}
}

// Direction graph
func (g *Graph) AddEdge(src, dest int) {
	newNode := NewAdjListNode(dest)
	newNode.Next = g.Array[src]
	g.Array[src] = newNode
}

func (g *Graph) BFS(s int) {
	visited := make(map[int]bool, g.V)

	q := queue.New(g.V)

	visited[s] = true
	q.Enqueue(s)

	for q.Len() != 0 {
		s, _ = q.Dequeue()
		fmt.Print(s, " ")

		node := g.Array[s]
		for node != nil {
			n := node.Vertex
			if !visited[n] {
				visited[n] = true
				q.Enqueue(n)
			}
			node = node.Next
		}
	}
}

func (g *Graph) String() string {
	var str strings.Builder
	var v int
	for v = 0; v < g.V; v++ {
		pCrawl := g.Array[v]
		str.WriteString(fmt.Sprintf("\n Ajacency list of vertex %d\n head ", v))
		for pCrawl != nil {
			str.WriteString(fmt.Sprintf("-> %d", pCrawl.Vertex))
			pCrawl = pCrawl.Next
		}
		str.WriteString("\n")
	}
	return str.String()
}
