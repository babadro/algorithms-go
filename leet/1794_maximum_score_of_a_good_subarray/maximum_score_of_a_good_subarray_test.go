package _1793_maximum_score_of_a_good_subarray

import (
	"fmt"
	"testing"
)

func Test_maximumScore(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 4, 3, 7, 4, 5}, 3, 15},
		{[]int{5, 5, 4, 5, 4, 1, 1, 1}, 0, 20},
		{[]int{5}, 0, 5},
		{[]int{1, 2, 5}, 2, 5},
		{[]int{6569, 9667, 3148, 7698, 1622, 2194, 793, 9041, 1670, 1872}, 5, 9732},
		{arr474, 474, 36808},     // todo 1 fails.
		{arr44424, 44424, 96844}, // todo 1 fails
	}
	for _, tt := range tests {
		var name string
		if len(tt.nums) > 5 {
			name = fmt.Sprintf("%v<...>", tt.nums[:5])
		} else {
			name = fmt.Sprintf("%v", tt.nums)
		}

		t.Run(name, func(t *testing.T) {
			if got := maximumScore(tt.nums, tt.k); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
