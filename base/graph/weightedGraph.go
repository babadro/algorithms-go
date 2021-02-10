package graph

// 1 todo https://www.geeksforgeeks.org/dijkstras-algorithm-for-adjacency-list-representation-greedy-algo-8/

type Edge struct {
	src, dst, weight int
}

type WeightedUndirectedGraph struct {
	V     int
	edges []Edge
}

func (w *WeightedUndirectedGraph) AddEdge(src, dst, weight int) {
	w.edges = append(w.edges, Edge{
		src:    src,
		dst:    dst,
		weight: weight,
	})
}
