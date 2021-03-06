package adjgraphBFS

import (
	"fmt"
	"github.com/babadro/algorithms-go/base/queue"
	"math"
	"strings"
)

type color int

const (
	white color = 0
	gray  color = 1
	black color = 2
)

type AdjListNode struct {
	Vertex int
	Next   *AdjListNode
}

func NewAdjListNode(dest int) *AdjListNode {
	return &AdjListNode{dest, nil}
}

type Vertex struct {
	Color       color
	Predecessor *Vertex
	Distance    int
}

type Graph struct {
	V        int
	Array    []*AdjListNode
	Vertices []Vertex
}

func CreateGraph(V int) *Graph {
	return &Graph{V: V, Array: make([]*AdjListNode, V), Vertices: make([]Vertex, V)}
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

func BFS(g *Graph) (source int) {
	for i, u := range g.Vertices {
		if i == source {
			continue
		}

		u.Color = white
		u.Distance = math.MaxInt64
		u.Predecessor = nil
	}

	s := g.Vertices[source]
	s.Color = gray
	s.Distance = 0
	s.Predecessor = nil

	q := queue.queue{}

}
