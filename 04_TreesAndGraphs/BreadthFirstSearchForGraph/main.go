package main

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

func (g *Graph) BFS(s int) {
	visited := make([]bool, g.V)

	var queue []int
	queue = append(queue, s)

	for len(queue) > 0 {
		s = queue[len(queue)-1]
		queue = queue[:(len(queue) - 1)]

		fmt.Println(s)
		visited[s] = true

		node := g.Array[s]
		for node != nil {
			if !visited[node.Vertex] {
				queue = append(queue, node.Vertex)
				visited[node.Vertex] = true
			}
			node = node.Next
		}
	}
}

func main() {
	g := CreateGraph(4)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	fmt.Println("Following is Breadth First Traversal \n (starting from vertex 2")

	g.BFS(2)
}
