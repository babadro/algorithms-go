package main

type AdjNode struct {
	Vertex int
	Next   int
}

type Graph []AdjNode

func NewGraph(vertexCount int) Graph {
	return make([]AdjNode, vertexCount)
}

func main() {
	graph := NewGraph(5)

}
