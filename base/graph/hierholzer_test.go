package graph

import (
	"reflect"
	"testing"
)

func TestEulerianCircuit(t *testing.T) {
	tests := []struct {
		edges [][]int
		n     int
		want  []int
	}{
		{
			[][]int{{0, 1}, {1, 2}, {2, 0}},
			3,
			[]int{0, 1, 2, 0},
		},
		{
			[][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {3, 4}, {4, 1}},
			5,
			[]int{0, 1, 3, 4, 1, 2, 0},
		},
		{
			[][]int{{0, 6}, {0, 1}, {1, 2}, {2, 0}, {2, 3}, {3, 4}, {4, 2}, {4, 5}, {5, 0}, {6, 4}},
			7,
			[]int{0, 1, 2, 3, 4, 5, 0, 6, 4, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			g := New(tt.n)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			if got := EulerianCircuit(g.adj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EulerianCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
