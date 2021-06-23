package _2_connecting_n_pipes_with_minimum_cost

import (
	"fmt"
	"testing"
)

func Test_connectSticks(t *testing.T) {
	tests := []struct {
		sticks []int
		want   int
	}{
		{[]int{4, 3, 2, 6}, 29},
		{[]int{2, 4, 3}, 14},
		{[]int{1, 8, 3, 5}, 30},
		{[]int{3}, 0},
		{[]int{4, 4, 4, 6}, 36},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.sticks), func(t *testing.T) {
			if got := connectSticks(tt.sticks); got != tt.want {
				t.Errorf("connectSticks() = %v, want %v", got, tt.want)
			}
		})
	}
}
