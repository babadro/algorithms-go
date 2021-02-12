package spanTree

import (
	"github.com/babadro/algorithms-go/base/unionFind"
	"math"
	"sort"
)

type Edge struct {
	src, dst, weight int
}

// Kruskal minimum spanning tree
func KruskalMST(vertexCount int, edges []Edge) (mst []Edge) {
	union := unionFind.NewWQUPC(vertexCount)

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

// https://www.programiz.com/dsa/prim-algorithm
func PrimMST(vertexCount int, adjMatrix [][]int) (mst []Edge) {
	edgesCount := 0
	selected := make([]bool, vertexCount)
	selected[0] = true

	for edgesCount < vertexCount-1 {
		min := math.MaxInt64
		x, y := 0, 0

		for i := 0; i < vertexCount; i++ {
			if selected[i] {
				for j := 0; j < vertexCount; j++ {
					if !selected[j] && adjMatrix[i][j] != 0 {
						if min > adjMatrix[i][j] {
							min = adjMatrix[i][j]
							x, y = i, j
						}
					}
				}
			}
		}

		mst = append(mst, Edge{
			src:    x,
			dst:    y,
			weight: adjMatrix[x][y],
		})

		selected[y] = true
		edgesCount++
	}

	return mst
}
