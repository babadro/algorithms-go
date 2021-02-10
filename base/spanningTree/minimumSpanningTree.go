package spanTree

import (
	"github.com/babadro/algorithms-go/base/unionFind"
	"sort"
)

type Edge struct {
	src, dst, weight int
}

// Kruskal minimum spanning tree
// V - vertex count for given graph
// todo 1 unit tests
func KruskalMST(V int, edges []Edge) (mst []Edge) {
	union := unionFind.NewWQUPC(V)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	for _, edge := range edges {
		if union.Find(edge.src, edge.dst) {
			continue
		}

		union.Union(edge.src, edge.dst)

		mst = append(mst, edge)
	}

	return mst
}
