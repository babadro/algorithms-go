package _2713_maximum_strictly_increasing_cells_in_a_matrix

import "testing"

func Test_maxIncreasingCells(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want int
	}{
		{
			mat: [][]int{
				{3, 1},
				{3, 4},
			},
			want: 2,
		},
		{
			mat: [][]int{
				{1, 1},
				{1, 1},
			},
			want: 1,
		},
		{
			mat: [][]int{
				{3, 1, 6},
				{-9, 5, 7},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIncreasingCells(tt.mat); got != tt.want {
				t.Errorf("maxIncreasingCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
