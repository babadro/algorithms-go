package _1727_largest_submatrix_with_rearrangements

import "testing"

func Test_largestSubmatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{
			"1",
			[][]int{
				{0, 0, 1},
				{1, 1, 1},
				{1, 0, 1},
			},
			4,
		},
		{
			"2",
			[][]int{
				{1, 0, 1, 0, 1},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSubmatrix(tt.matrix); got != tt.want {
				t.Errorf("largestSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
