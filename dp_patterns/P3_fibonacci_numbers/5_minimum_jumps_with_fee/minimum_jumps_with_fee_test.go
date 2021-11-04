package __minimum_jumps_with_fee

import (
	"fmt"
	"testing"
)

func Test_findMinFee(t *testing.T) {
	tests := []struct {
		fee  []int
		want int
	}{
		{[]int{1, 2, 5, 2, 1, 2}, 3},
		{[]int{2, 3, 4, 5}, 5},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.fee), func(t *testing.T) {
			if got := findMinFeeOptimized(tt.fee); got != tt.want {
				t.Errorf("findMinFee() = %v, want %v", got, tt.want)
			}
		})
	}
}
