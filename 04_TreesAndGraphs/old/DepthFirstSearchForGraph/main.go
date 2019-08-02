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

// Direction graph
func (g *Graph) AddEdge(src, dest int) {
	newNode := NewAdjListNode(dest)
	newNode.Next = g.Array[src].Head
	g.Array[src].Head = newNode
}

func (g *Graph) String() string {
	var str strings.Builder
	var v int
	for v = 0; v < g.V; v++ {
		pCrawl := g.Array[v].Head
		str.WriteString(fmt.Sprintf("\n Ajacency list of vertex %d\n head ", v))
		for pCrawl != nil {
			str.WriteString(fmt.Sprintf("-> %d", pCrawl.Vertex))
			pCrawl = pCrawl.Next
		}
		str.WriteString("\n")
	}
	return str.String()
}

func (g *Graph) dFSUtil(vertex int, visited []bool) {
	visited[vertex] = true
	fmt.Println(vertex)

	node := g.Array[vertex].Head

	for node != nil {
		if !visited[node.Vertex] {
			g.dFSUtil(node.Vertex, visited)
		}
		node = node.Next
	}
}

func (g *Graph) DFS() {
	visited := make([]bool, g.V)

	for i := 0; i < g.V; i++ {
		if !visited[i] {
			g.dFSUtil(i, visited)
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

	//fmt.Println(g)

	fmt.Println("Following is Depth First Traversal")

	g.DFS()
}
