package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
		want int
	}{
		{0, []int{0}, 0},
		{1, nil, -1},
		{2, []int{1, 2, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %d", tt.nums, tt.num), func(t *testing.T) {
			if got := BinarySearch(tt.num, tt.nums); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
