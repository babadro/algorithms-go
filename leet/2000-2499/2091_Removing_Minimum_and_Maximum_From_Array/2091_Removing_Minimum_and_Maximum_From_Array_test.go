package Removing_Minimum_and_Maximum_From_Array

import (
	"fmt"
	"testing"
)

func Test_minimumDeletions(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 10, 7, 5, 4, 1, 8, 6}, 5},
		{[]int{0, -4, 19, 1, 8, -2, -3, 5}, 3},
		{[]int{101}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := minimumDeletions(tt.nums); got != tt.want {
				t.Errorf("minimumDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}
