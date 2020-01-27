package _1334_find_city_with_smallest_number_neighbors_and_distance

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// create graph
	graph := CreateGraph(n)
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1], edge[2])
	}

	var res *AdjListNode
	visited := make(map[*AdjListNode]bool)
	for i, list := range graph.Array {
		node := list.Head
		for node != nil {
			for k := range visited {
				delete(visited, k)
			}
			graph.visit(node, visited, distanceThreshold)
			node = node.Next
		}
	}
}

func (g *Graph) visit(node *AdjListNode, visited map[*AdjListNode]bool, budget int) {
	if node.Weight > budget || visited[node] {
		return
	}
	visited[node] = true
	budget -= node.Weight
	node := g.Array[node.Vertex].Head
	for node != nil {
		if !visited[node]
	}
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

func