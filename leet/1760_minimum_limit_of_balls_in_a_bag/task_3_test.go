package _1760_minimum_limit_of_balls_in_a_bag

import (
	"fmt"
	"testing"
)

func Test_minimumSize(t *testing.T) {
	tests := []struct {
		nums          []int
		maxOperations int
		want          int
	}{
		{
			[]int{9},
			2,
			3,
		},
		{
			[]int{2, 4, 8, 2},
			4,
			2,
		},
		{
			[]int{7, 17},
			2,
			7,
		},
		{
			[]int{1},
			1,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %d", tt.nums, tt.maxOperations), func(t *testing.T) {
			if got := minimumSize(tt.nums, tt.maxOperations); got != tt.want {
				t.Errorf("minimumSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
