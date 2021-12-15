package graph

import (
	"testing"
)

func TestEulerianPathGraph(t *testing.T) {
	tests := []struct {
		edges [][2]int
		n     int
		want  bool
	}{
		{[][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}, {1, 4}}, 5, true},
		{[][2]int{{4, 3}, {3, 0}, {0, 5}, {5, 4}}, 6, true},
		{[][2]int{{0, 1}, {1, 2}, {2, 0}}, 3, true},
		{[][2]int{{0, 1}, {1, 2}, {2, 0}, {2, 3}, {3, 0}}, 4, true},
		{[][2]int{{1, 0}, {0, 2}, {2, 1}, {1, 3}, {3, 0}, {3, 4}}, 5, false},
		{[][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}, {1, 4}, {4, 3}, {3, 0}, {0, 5}, {5, 4}}, 6, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			g := NewEulerianPathGraph(tt.n)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			if got := HasEulerPath(*g); got != tt.want {
				t.Errorf("NewEulerianPathGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
