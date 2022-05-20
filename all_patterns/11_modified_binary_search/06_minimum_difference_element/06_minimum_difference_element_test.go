package _6_minimum_difference_element

import (
	"fmt"
	"testing"
)

func Test_searchMinDiffElement(t *testing.T) {
	tests := []struct {
		arr  []int
		key  int
		want int
	}{
		{[]int{4, 6, 10}, 7, 6},
		{[]int{4, 6, 10}, 4, 4},
		{[]int{1, 3, 8, 10, 15}, 12, 10},
		{[]int{4, 6, 10}, 17, 10},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			if got := searchMinDiffElement(tt.arr, tt.key); got != tt.want {
				t.Errorf("searchMinDiffElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
