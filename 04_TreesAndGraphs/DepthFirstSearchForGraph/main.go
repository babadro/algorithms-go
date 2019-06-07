package main

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

func (g *Graph) AddEdge(src, dest int) {
	newNode := NewAdjListNode(dest)
	newNode.Next = g.Array[src].Head
	g.Array[src].Head = newNode

	newNode = NewAdjListNode(src)
	newNode.Next = g.Array[dest].Head
	g.Array[dest].Head = newNode
}
