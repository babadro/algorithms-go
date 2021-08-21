package _1049_last_stone_weight_ii

import (
	"fmt"
	"testing"
)

func Test_lastStoneWeightII(t *testing.T) {
	tests := []struct {
		stones []int
		want   int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{31, 26, 33, 21, 40}, 5},
		{[]int{1, 2}, 1},
		{[]int{1}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.stones), func(t *testing.T) {
			if got := lastStoneWeightII(tt.stones); got != tt.want {
				t.Errorf("lastStoneWeightII() = %v, want %v", got, tt.want)
			}
		})
	}
}
