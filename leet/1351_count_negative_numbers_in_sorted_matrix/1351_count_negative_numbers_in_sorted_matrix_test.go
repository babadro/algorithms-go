package _1351_count_negative_numbers_in_sorted_matrix

import "testing"

func Test_countNegatives(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			"1",
			[][]int{
				{4, 3, 2, -1},
				{3, 2, 1, -1},
				{1, 1, -1, -2},
				{-1, -1, -2, -3},
			},
			8,
		},
		{
			"2",
			[][]int{
				{3, 2},
				{1, 0},
			},
			0,
		},
		{
			"3",
			[][]int{
				{-1},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNegatives2(tt.grid); got != tt.want {
				t.Errorf("countNegatives() = %v, want %v", got, tt.want)
			}
		})
	}
}
