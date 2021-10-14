package count_of_subset_sum

import (
	"fmt"
	"testing"
)

func Test_countSubset(t *testing.T) {
	tests := []struct {
		nums []int
		sum  int
		want int
	}{
		{[]int{1, 1, 2, 3}, 4, 3},
		{[]int{1, 2, 7, 1, 5}, 9, 3},
		{[]int{1, 1, 1}, 1, 3},
		{[]int{1, 1, 1}, 3, 1},
		{[]int{1, 1, 1}, 2, 3},
		{[]int{1, 1, 1, 1}, 2, 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.nums, tt.sum), func(t *testing.T) {
			if got := countSubsetBottomUp(tt.nums, tt.sum); got != tt.want {
				t.Errorf("countSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
