package _2863_maximum_lengh_of_semi_decreasing_subarrays

import "testing"

func Test_maxSubarrayLength(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 0},
		{[]int{7, 6, 5, 4, 3, 2, 1, 10, 11}, 7},
		{[]int{7, 6, 5, 4, 3, 2, 1, 6, 10, 11}, 8},
		{[]int{57, 55, 50, 60, 61, 58, 63, 59, 64, 60, 63}, 6},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxSubarrayLength(tt.nums); got != tt.want {
				t.Errorf("maxSubarrayLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
