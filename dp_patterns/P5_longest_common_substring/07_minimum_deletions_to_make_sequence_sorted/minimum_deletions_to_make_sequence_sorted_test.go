package _7_minimum_deletions_to_make_sequence_sorted

import (
	"fmt"
	"testing"
)

func Test_findMinDeletions(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{4, 2, 3, 6, 10, 1, 12}, 2},
		{[]int{-4, 10, 3, 7, 15}, 1},
		{[]int{3, 2, 1, 0}, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := findMinDeletions(tt.nums); got != tt.want {
				t.Errorf("findMinDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}
