package _033_search_in_rotated_sorted_array

import (
	"fmt"
	"testing"
)

func TestSearchNotUnique(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{1}, 0, -1},
		{[]int{1}, 1, 0},
		{[]int{2, 1}, 1, 1},
		{[]int{3, 1, 2}, 2, 2},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1, 3}, 1, 0},
		{[]int{1, 2, 3, 4}, 2, 1},
		{[]int{3, 1}, 3, 0},
		{[]int{3, 1, 2}, 3, 0},

		{[]int{3, 4, 5, 1, 2, 3}, 3, 0},
		{[]int{4, 4, 4, 4, 4}, 4, 0},
		{[]int{4, 4, 4, 1, 2, 3}, 4, 0},
		{[]int{4, 4, 4, 1, 2, 3, 4}, 4, 0},
		{[]int{4, 5, 6, 6, 4, 4, 4}, 4, 0},
		{[]int{4, 5, 6, 3, 3, 3, 3}, 3, 3},
	}

	for i, c := range cases {
		if fact := searchNotUnique(c.nums, c.target); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}

func Test_binarySearchTarget(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{0, 1, 2, 4, 5, 6, 7}, 4, 3},
		{[]int{1}, 0, -1},
		{[]int{1}, 1, 0},
		{[]int{1, 2}, 1, 0},
		{[]int{1, 2, 3}, 2, 1},
		{[]int{0, 1, 2, 4, 5, 6, 7}, 3, -1},
		{[]int{1, 3}, 1, 0},
		{[]int{1, 2, 3, 4}, 2, 1},
		{[]int{1, 2, 2, 2, 3}, 2, 1},
		{[]int{3, 3, 3, 3}, 3, 0},
		{[]int{1, 2, 3, 3, 3, 3}, 3, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := binarySearchTarget(tt.nums, 0, len(tt.nums), tt.target); got != tt.want {
				t.Errorf("binarySearchTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
