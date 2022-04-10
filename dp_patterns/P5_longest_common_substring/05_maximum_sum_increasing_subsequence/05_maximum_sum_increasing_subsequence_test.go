package _5_maximum_sum_increasing_subsequence

import (
	"fmt"
	"testing"
)

func Test_findMSISBruteForce(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{4, 1, 2, 6, 10, 1, 12}, 32},
		{[]int{-4, 10, 3, 7, 15}, 25},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := findMSISBottomUp(tt.nums); got != tt.want {
				t.Errorf("findMSIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
