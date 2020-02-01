package _1334_find_city_with_smallest_number_neighbors_and_distance

import (
	"math"
)

func findTheCity(n int, edges [][3]int, distanceThreshold int) int {
	graph := CreateGraph(n)
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1], edge[2])
	}
	minNumReachableCities := math.MaxInt32
	maxVertex := -1
	for vertex := range graph.Array {
		numReachableCities := graph.checkVertex(vertex, distanceThreshold)

		if numReachableCities < minNumReachableCities || (numReachableCities == minNumReachableCities && vertex > maxVertex) {
			minNumReachableCities, maxVertex = numReachableCities, vertex
		}
	}
	return maxVertex
}

func (g *Graph) checkVertex(vertex, threshold int) int {
	node := g.Array[vertex].Head
	visited := make(map[int]bool)
	for node != nil {
		g.visit(node, visited, threshold, vertex)
		node = node.Next
	}
	return len(visited)
}

func (g *Graph) visit(node *AdjListNode, visited map[int]bool, threshold int, startFrom int) {
	if node.Vertex == startFrom {
		return
	}
	if visited[node.Vertex] {
		return
	}
	threshold = threshold - node.Weight
	if threshold < 0 {
		return
	}
	visited[node.Vertex] = true
	node = g.Array[node.Vertex].Head
	for node != nil {
		g.visit(node, visited, threshold, startFrom)
		node = node.Next
	}
	return
}

type AdjListNode struct {
	Vertex int
	Weight int
	Next   *AdjListNode
}

func NewAdjListNode(dest, weight int) *AdjListNode {
	return &AdjListNode{dest, weight, nil}
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

func (g *Graph) AddEdge(src, dest, weight int) {
	newNode := NewAdjListNode(dest, weight)
	newNode.Next = g.Array[src].Head
	g.Array[src].Head = newNode

	newNode = NewAdjListNode(src, weight)
	newNode.Next = g.Array[dest].Head
	g.Array[dest].Head = newNode
}
