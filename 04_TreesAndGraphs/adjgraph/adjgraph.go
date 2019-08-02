package adjgraph

import (
	"fmt"
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

func CreateGraph(V int) *Graph {
	return &Graph{V, make([]*AdjListNode, V)}
}

// Direction graph
func (g *Graph) AddEdge(src, dest int) {
	newNode := NewAdjListNode(dest)
	newNode.Next = g.Array[src]
	g.Array[src] = newNode
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
