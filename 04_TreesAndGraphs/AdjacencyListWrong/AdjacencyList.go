package main

import (
	"fmt"
	"strings"
)

type AdjNode struct {
	Vertex int
	Next   *AdjNode
}

func NewAdjNode(vertex int) AdjNode {
	return AdjNode{vertex, nil}
}

type Graph struct {
	Data []AdjNode
}

func NewGraph(vertexCount int) *Graph {
	return &Graph{make([]AdjNode, vertexCount)}
}

func (g *Graph) AddEdge(src int, dst int) {
	node := NewAdjNode(dst)
	node.Next = &g.Data[src]
	g.Data[src] = node

	node = NewAdjNode(src)
	node.Next = &g.Data[dst]
	g.Data[dst] = node
}

// Doesn't work. Infinitive loop
func (g *Graph) String() string {
	var str strings.Builder

	for i := range g.Data {
		str.WriteString(fmt.Sprintf("Adjacency list of vertex %v\n head", i))
		temp := &g.Data[i]
		for temp != nil {
			str.WriteString(fmt.Sprintf(" -> %v", temp.Vertex))
			temp = temp.Next
		}
		str.WriteString(" \n")
	}
	return str.String()
}

func main() {
	graph := NewGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	fmt.Println(graph)
}
