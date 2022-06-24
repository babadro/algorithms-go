package _931_minimum_failing_path_sum

import "testing"

func Test_minFallingPathSum(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		{[][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}, 13},
		{tleInput1, -1428},
		{[][]int{{-48}}, -48},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minFallingPathSum2(tt.matrix); got != tt.want {
				t.Errorf("minFallingPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
