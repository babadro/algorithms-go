package _2316_count_Unreachable_pairs_of_node_in_undirected_graph

import "testing"

func Test_countPairs(t *testing.T) {
	tests := []struct {
		n     int
		edges [][]int
		want  int64
	}{
		{3, [][]int{{0, 1}, {0, 2}, {1, 2}}, 0},
		{
			7, [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}, 14,
		},
		{
			100000, [][]int{}, 4999950000,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := countPairs(tt.n, tt.edges); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
