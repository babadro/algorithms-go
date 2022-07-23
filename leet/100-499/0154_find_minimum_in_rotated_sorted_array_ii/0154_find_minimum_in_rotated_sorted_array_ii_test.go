package _154_find_minimum_in_rotated_sorted_array_ii

import (
	"fmt"
	"testing"
)

func Test_findMin(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		// unique cases
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{0, 1, 2, 4, 5, 6, 7}, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
		{[]int{1}, 1},
		//// cases with duplicates
		{[]int{2, 2, 2, 0, 1}, 0},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 1},
		{[]int{2, 2, 2, 3, 2, 2, 2}, 2},
		{[]int{3, 1, 3}, 1},
		{[]int{10, 1, 10, 10, 10}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := findMin(tt.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
