package _334_increasing_triplet_subsequence

import (
	"fmt"
	"testing"
)

func Test_increasingTriplet(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{1, 2}, false},
		{[]int{1, 2, 3}, true},
		{[]int{1, 2, 2}, false},
		{[]int{1, 2, 2, 4, 5}, true},
		{[]int{5, 1, 5, 5, 2, 5, 4}, true},
		{[]int{1, 1, -2, 6}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := increasingTriplet(tt.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
