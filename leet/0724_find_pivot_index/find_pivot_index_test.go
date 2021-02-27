package _724_find_pivot_index

import (
	"fmt"
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
		{[]int{1}, 0},
		{[]int{0}, 0},
		{[]int{1, -1}, -1},
		{[]int{1, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			if got := pivotIndex(tt.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
