package _081_search_in_rotated_sorted_array_ii

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   bool
	}{
		{[]int{1}, 1, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{1, 1, 1, 1, 1, 1, 1}, 1, true},
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{1, 2, 3}, 0, false},
		{[]int{1, 2, 3, 3, 3, 3, 4, 5}, 1, true},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2, true},
		{[]int{2, 2, 2, 3, 2, 2, 2}, 3, true},
		{[]int{3, 1}, 3, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
