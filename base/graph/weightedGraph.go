package graph

// 1 todo https://www.geeksforgeeks.org/dijkstras-algorithm-for-adjacency-list-representation-greedy-algo-8/

type node struct {
	dst    int
	weight int
}

type WeightedGraph struct {
	directed bool
	adg      [][]node
}

func NewWeighted(v int, directed bool) *WeightedGraph {
	return &WeightedGraph{directed, make([][]node, v)}
}

func (w *WeightedGraph) AddEdge(src, dst, weight int) {
	w.adg[src] = append(w.adg[src], node{dst, weight})
	if !w.directed {
		w.adg[dst] = append(w.adg[dst], node{src, weight})
	}
}
