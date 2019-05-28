package main

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
	Array *AdjList
}

func CreateGraph(V int) *Graph {
	g := &Graph{V: V}
	g.Array = &AdjList{}
}
