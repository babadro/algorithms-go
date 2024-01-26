package _490_the_mase

import "testing"

func Test_hasPath(t *testing.T) {
	tests := []struct {
		maze        [][]int
		start       []int
		destination []int
		want        bool
	}{
		{
			[][]int{{0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}, {1, 1, 0, 1, 1}, {0, 0, 0, 0, 0}},
			[]int{0, 4},
			[]int{4, 4},
			true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := hasPath(tt.maze, tt.start, tt.destination); got != tt.want {
				t.Errorf("hasPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
