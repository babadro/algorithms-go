package _070_climbing_stairs

import (
	"strconv"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{2, 2},
		{1, 1},
		{3, 3},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			if got := climbStairsTopDown(tt.n); got != tt.want {
				t.Errorf("climbStairs2() = %v, want %v", got, tt.want)
			}
		})
	}
}
