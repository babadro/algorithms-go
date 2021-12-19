package graph

import (
	"fmt"
	"testing"
)

func TestEulerianGraph_HasEulerianCycle(t *testing.T) {
	tests := []struct {
		edges [][2]int
		n     int
		want  bool
	}{
		{
			[][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}, {1, 4}, {4, 3}, {3, 0}}, 5, true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.edges), func(t *testing.T) {
			g := NewEulerianGraph(tt.n)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			if got := g.HasEulerianCycle(); got != tt.want {
				t.Errorf("HasEulerianCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
