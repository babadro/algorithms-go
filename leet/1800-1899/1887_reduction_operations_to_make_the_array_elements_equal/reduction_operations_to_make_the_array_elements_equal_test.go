package _1887_reduction_operations_to_make_the_array_elements_equal

import (
	"fmt"
	"testing"
)

func Test_reductionOperations(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{5, 1, 3}, 3},
		{[]int{1, 1, 1}, 0},
		{[]int{1, 1, 2, 2, 3}, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := reductionOperations2(tt.nums); got != tt.want {
				t.Errorf("reductionOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
