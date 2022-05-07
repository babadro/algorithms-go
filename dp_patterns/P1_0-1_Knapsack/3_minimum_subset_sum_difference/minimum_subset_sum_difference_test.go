package __minimum_subset_sum_difference

import (
	"fmt"
	"testing"
)

func Test_minDiff(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 9}, 3},
		{[]int{1, 2, 7, 1, 5}, 0},
		{[]int{1, 3, 100, 4}, 92},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := minDiffBottomUp(tt.nums); got != tt.want {
				t.Errorf("minDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
