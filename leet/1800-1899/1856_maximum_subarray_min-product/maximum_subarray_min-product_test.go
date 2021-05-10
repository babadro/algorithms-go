package _1856_maximum_subarray_min_product

import (
	"fmt"
	"testing"
)

func Test_maxSumMinProduct(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 2}, 14},
		{[]int{2, 3, 3, 1, 2}, 18},
		{[]int{3, 1, 5, 6, 4, 2}, 60},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := maxSumMinProduct(tt.nums); got != tt.want {
				t.Errorf("maxSumMinProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
