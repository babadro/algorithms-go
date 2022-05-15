package _6_triplets_with_smaller_sum

import (
	"fmt"
	"testing"
)

func Test_searchTriplets(t *testing.T) {
	tests := []struct {
		arr     []int
		target  int
		wantRes int
	}{
		{[]int{-1, 0, 2, 3}, 3, 2},
		{[]int{-1, 4, 2, 1, 3}, 5, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			if gotRes := searchTriplets(tt.arr, tt.target); gotRes != tt.wantRes {
				t.Errorf("searchTriplets() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
