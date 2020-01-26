package _1334_find_city_with_smallest_number_neighbors_and_distance

func findTheCity(n int, edges [][]int, distanceThreshold int) int {

}

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
