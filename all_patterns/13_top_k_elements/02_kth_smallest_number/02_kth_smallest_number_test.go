package _2_kth_smallest_number

import "testing"

func Test_findKthSmallestNumber(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 5, 12, 2, 11, 5}, 3, 5},
		{[]int{1, 5, 12, 2, 11, 5}, 4, 5},
		{[]int{5, 12, 11, -1, 12}, 3, 11},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findKthSmallestNumber(tt.nums, tt.k); got != tt.want {
				t.Errorf("findKthSmallestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
