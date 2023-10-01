package _2088_count_fertile_piramyds_in_a_land

import "testing"

func Test_countPyramids(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 1, 1, 0}, {1, 1, 1, 1}}, 2},
		{[][]int{{1, 1, 1}, {1, 1, 1}}, 2},
		{[][]int{{1, 1, 1, 1, 0}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {0, 1, 0, 0, 1}}, 13},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := countPyramids(tt.grid); got != tt.want {
				t.Errorf("countPyramids() = %v, want %v", got, tt.want)
			}
		})
	}
}
