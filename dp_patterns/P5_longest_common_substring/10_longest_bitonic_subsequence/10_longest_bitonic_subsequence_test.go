package _10_longest_bitonic_subsequence

import (
	"fmt"
	"testing"
)

func Test_findLBSLenBruteForce(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{4, 2, 3, 6, 10, 1, 12}, 5},
		{[]int{4, 2, 5, 9, 7, 6, 10, 3, 1}, 7},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := findLBSLenTopDown(tt.nums); got != tt.want {
				t.Errorf("findLBSLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
