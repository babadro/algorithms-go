package _1_order_agnostic_binary_search

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	tests := []struct {
		arr  []int
		key  int
		want int
	}{
		{[]int{4, 6, 10}, 10, 2},
		{[]int{1}, 1, 0},
		{[]int{0}, 1, -1},
		{[]int{1, 2, 3}, 1, 0},
		// {[]int{5, 5, 4, 4, 3, 3, 2, 1}, 3, 4}, returns 5
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.arr, tt.key), func(t *testing.T) {
			if got := search(tt.arr, tt.key); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
