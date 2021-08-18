package _416_partition_equal_subset_sum

import (
	"fmt"
	"testing"
)

func Test_canPartition(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		//{[]int{1, 5, 11, 5}, true},
		//{[]int{1, 2, 3, 5}, false},
		//{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, true},
		//{[]int{1, 1, 2, 2, 3, 3, 4, 4}, true},
		{tleInput1, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := canPartition(tt.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
