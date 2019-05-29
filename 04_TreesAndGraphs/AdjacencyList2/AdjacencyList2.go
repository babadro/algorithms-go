package main

import (
	"fmt"
	"strings"
)

type AdjListNode struct {
	Dest int
	Next *AdjListNode
}

func NewAdjListNode(dest int) *AdjListNode {
	return &AdjListNode{dest, nil}
}

type AdjList struct {
	Head *AdjListNode
}

type Graph struct {
	V     int
	Array []AdjList
}

func CreateGraph(V int) *Graph {
	return &Graph{V, make([]AdjList, V)}
}

func (g *Graph) AddEdge(src, dest int) {
	newNode := NewAdjListNode(dest)
	newNode.Next = g.Array[src].Head
	g.Array[src].Head = newNode

	newNode = NewAdjListNode(src)
	newNode.Next = g.Array[dest].Head
	g.Array[dest].Head = newNode
}

func (g *Graph) String() string {
	var str strings.Builder
	var v int
	for v = 0; v < g.V; v++ {
		pCrawl := g.Array[v].Head
		str.WriteString(fmt.Sprintf("\n Ajacency list of vertex %d\n head ", v))
		for pCrawl != nil {
			str.WriteString(fmt.Sprintf("-> %d", pCrawl.Dest))
			pCrawl = pCrawl.Next
		}
		str.WriteString("\n")
	}
	return str.String()
}

func main() {
	V := 5
	graph := CreateGraph(V)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	fmt.Println(graph)
}
