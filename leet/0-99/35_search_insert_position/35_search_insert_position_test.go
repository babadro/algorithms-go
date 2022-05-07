package _5_search_insert_position

import (
	"fmt"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	tests := []struct {
		arr  []int
		key  int
		want int
	}{
		{[]int{4, 6, 10}, 6, 1},
		{[]int{1, 3, 8, 10, 15}, 12, 4},
		{[]int{4, 6, 10}, 17, 3},
		{[]int{4, 6, 10}, -1, 0},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			if got := searchInsert(tt.arr, tt.key); got != tt.want {
				t.Errorf("searchCeilingOfANumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
