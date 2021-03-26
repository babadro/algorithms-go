package _1803_count_pairs_with_xor_in_a_range

import (
	"fmt"
	"testing"
)

func Test_countPairs(t *testing.T) {
	tests := []struct {
		nums []int
		low  int
		high int
		want int
	}{
		{[]int{1, 4, 2, 7}, 2, 6, 6},
		{[]int{9, 8, 4, 2, 1}, 5, 14, 8},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := countPairs(tt.nums, tt.low, tt.high); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
