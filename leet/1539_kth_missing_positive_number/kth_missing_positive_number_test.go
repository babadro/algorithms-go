package _1539_kth_missing_positive_number

import (
	"fmt"
	"testing"
)

func Test_findKthPositive(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		want int
	}{
		{[]int{2, 3, 4, 7, 11}, 5, 9},
		{[]int{1, 2, 3, 4}, 2, 6},
		{[]int{1}, 1, 2},
		{[]int{3, 4}, 1, 1},
		{[]int{3, 4}, 2, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			if got := findKthPositive2(tt.arr, tt.k); got != tt.want {
				t.Errorf("findKthPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
