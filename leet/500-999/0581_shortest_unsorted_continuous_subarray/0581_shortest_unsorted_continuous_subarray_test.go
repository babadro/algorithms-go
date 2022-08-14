package _0581_shortest_unsorted_continuous_subarray

import (
	"fmt"
	"testing"
)

func Test_findUnsortedSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
		{[]int{1, 2, 3}, 0},
		{[]int{1}, 0},
		{[]int{1, 2, 5, 3, 7, 10, 9, 12}, 5},
		{[]int{1, 3, 2, 0, -1, 7, 10}, 5},
		{[]int{3, 2, 1}, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := findUnsortedSubarray(tt.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
