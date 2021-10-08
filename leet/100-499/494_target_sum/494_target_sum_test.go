package _494_target_sum

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{1}, 1, 1},
		{[]int{2, 1, 1}, 2, 2},
		{[]int{1, 0}, 1, 2},
		{[]int{1, 2, 1}, 0, 2},
		{[]int{0, 0, 0}, 0, 8},
		{[]int{0, 0}, 0, 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findTargetSumBottomUpOptimized(tt.nums, tt.target); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
