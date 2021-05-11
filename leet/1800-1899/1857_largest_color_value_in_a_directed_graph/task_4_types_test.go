package _1857_largest_color_value_in_a_directed_graph

import (
	"fmt"
	"testing"
)

func TestGraph_IsCyclic(t *testing.T) {
	tests := []struct {
		vertexCount int
		edges       [][]int
		want        bool
	}{
		{4, [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 0}, {2, 3}, {3, 3}}, true},
		{4, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}, true},
		{1, [][]int{{0, 0}}, true},
		{10, [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}, {3, 5}, {5, 6}, {2, 7}, {6, 7}, {7, 8}, {3, 8}, {5, 8}, {8, 9}, {3, 9}, {6, 9}}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.edges), func(t *testing.T) {
			g := New(tt.vertexCount)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			if got := g.IsCyclic(); got != tt.want {
				t.Errorf("IsCyclic() = %v, want %v", got, tt.want)
			}
		})
	}
}
