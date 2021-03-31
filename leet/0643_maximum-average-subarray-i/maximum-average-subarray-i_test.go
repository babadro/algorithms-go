package _643_maximum_average_subarray_i

import (
	"fmt"
	"testing"
)

func Test_findMaxAverage(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{1}, 1, 1},
		{[]int{0}, 1, 0},
		{[]int{1, 2}, 1, 2},
		{[]int{0, 1, 100, 50, 0, 1}, 2, 75},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := findMaxAverage(tt.nums, tt.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
