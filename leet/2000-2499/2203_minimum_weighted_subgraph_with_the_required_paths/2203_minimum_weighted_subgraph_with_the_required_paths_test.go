package _2203_minimum_weighted_subgraph_with_the_required_paths

import (
	"testing"
)

func Test_minimumWeight(t *testing.T) {
	tests := []struct {
		n     int
		edges [][]int
		src1  int
		src2  int
		dest  int
		want  int64
	}{
		{n: 6, edges: [][]int{{0, 2, 2}, {0, 5, 6}, {1, 0, 3}, {1, 4, 5}, {2, 1, 1}, {2, 3, 3}, {2, 3, 4}, {3, 4, 2}, {4, 5, 1}},
			src1: 0, src2: 1, dest: 5, want: 9},
		{n: 3, edges: [][]int{{0, 1, 1}, {2, 1, 1}}, src1: 0, src2: 1, dest: 2, want: -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minimumWeight(tt.n, tt.edges, tt.src1, tt.src2, tt.dest); got != tt.want {
				t.Errorf("minimumWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
