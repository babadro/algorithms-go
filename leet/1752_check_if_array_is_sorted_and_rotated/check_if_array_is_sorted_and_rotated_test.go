package _1752_check_if_array_is_sorted_and_rotated

import (
	"fmt"
	"testing"
)

func Test_check2(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{3, 4, 5, 1, 2}, true},
		{[]int{2, 1, 3, 4}, false},
		{[]int{1, 2, 3}, true},
		{[]int{1, 1, 1}, true},
		{[]int{1}, true},
		{[]int{1, 2}, true},
		{[]int{2, 1}, true},
		{[]int{2, 1, 0}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			if got := check2(tt.nums); got != tt.want {
				t.Errorf("check2() = %v, want %v", got, tt.want)
			}
		})
	}
}
