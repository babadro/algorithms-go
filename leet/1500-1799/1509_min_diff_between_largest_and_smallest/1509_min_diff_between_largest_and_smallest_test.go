package _1509_min_diff_between_largest_and_smallest

import "testing"

func Test_minDifference(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{5, 3, 2, 4}, 0},
		{[]int{0, 1, 5, 10, 14}, 1},
		{[]int{6, 6, 0, 1, 1, 4, 6}, 2},
		{[]int{3, 100, 20}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minDifference(tt.nums); got != tt.want {
				t.Errorf("minDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
