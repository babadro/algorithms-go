package _4_longest_increasing_subsequence

import (
	"fmt"
	"testing"
)

func Test_findLISLengthBruteForce(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{4, 2, 3, 6, 10, 1, 12}, 5},
		{[]int{-4, 10, 3, 7, 15}, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := FindLISLenBottomUp(tt.nums); got != tt.want {
				t.Errorf("findLISLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
