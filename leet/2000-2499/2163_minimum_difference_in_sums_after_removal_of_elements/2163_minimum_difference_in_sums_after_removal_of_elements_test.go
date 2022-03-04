package _2163_minimum_difference_in_sums_after_removal_of_elements

import (
	"fmt"
	"testing"
)

func Test_minimumDifference(t *testing.T) {
	tests := []struct {
		nums []int
		want int64
	}{
		{[]int{3, 1, 2}, -1},
		{[]int{7, 9, 5, 8, 1, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := minimumDifference(tt.nums); got != tt.want {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
