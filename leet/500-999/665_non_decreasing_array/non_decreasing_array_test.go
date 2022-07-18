package _665_non_decreasing_array

import (
	"fmt"
	"testing"
)

func Test_checkPossibility(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		//{[]int{4, 2, 3}, true},
		//{[]int{4, 2, 1}, false},
		//{[]int{3, 4, 2, 3}, false},
		//{[]int{2, 4, 2, 2}, true},
		//{[]int{1}, true},
		//{[]int{5, 7, 1, 8}, true},
		{[]int{5, 4, 3, 10}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := checkPossibility(tt.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
