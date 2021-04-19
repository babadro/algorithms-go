package _861_score_after_flipping_matrix

import "testing"

func Test_matrixScore(t *testing.T) {
	tests := []struct {
		name string
		A    [][]int
		want int
	}{
		{
			A: [][]int{
				{0, 0, 1, 1},
				{1, 0, 1, 0},
				{1, 1, 0, 0},
			},
			want: 39,
		},
		{
			A: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{1, 1, 1},
			},
			want: 21,
		},
		{
			A: [][]int{
				{0, 1, 1},
				{0, 1, 1},
				{1, 1, 1},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixScore(tt.A); got != tt.want {
				t.Errorf("matrixScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
