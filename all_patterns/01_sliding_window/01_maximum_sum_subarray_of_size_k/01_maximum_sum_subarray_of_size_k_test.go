package _1_maximum_sum_subarray_of_size_k

import "testing"

func Test_findMaxSumSubArray(t *testing.T) {
	tests := []struct {
		k    int
		arr  []int
		want int
	}{
		{3, []int{2, 1, 5, 1, 3, 2}, 9},
		{2, []int{2, 3, 4, 1, 5}, 7},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findMaxSumSubArray(tt.k, tt.arr); got != tt.want {
				t.Errorf("findMaxSumSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
