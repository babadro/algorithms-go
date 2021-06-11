package _1886_determine_whether_matrix_can_be_obtained_by_rotation

import "testing"

func Test_findRotation(t *testing.T) {
	tests := []struct {
		mat    [][]int
		target [][]int
		want   bool
	}{
		{
			mat: [][]int{
				{0, 1},
				{1, 0},
			},
			target: [][]int{
				{1, 0},
				{0, 1},
			},
			want: true,
		},
		{
			mat: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
			target: [][]int{
				{1, 1, 1},
				{0, 1, 0},
				{0, 0, 0},
			},
			want: true,
		},
		{
			mat: [][]int{
				{0, 1},
				{1, 1},
			},
			target: [][]int{
				{1, 0},
				{0, 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findRotation(tt.mat, tt.target); got != tt.want {
				t.Errorf("findRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
