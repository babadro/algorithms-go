package _2_smallest_subarray_with_a_greater_sum

import (
	"fmt"
	"testing"
)

func Test_findMinSubArray(t *testing.T) {
	tests := []struct {
		arr  []int
		s    int
		want int
	}{
		{[]int{2, 1, 5, 2, 3, 2}, 7, 2},
		{[]int{2, 1, 5, 2, 8}, 7, 1},
		{[]int{3, 4, 1, 1, 6}, 8, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.arr, tt.s), func(t *testing.T) {
			if got := findMinSubArray(tt.s, tt.arr); got != tt.want {
				t.Errorf("findMinSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
