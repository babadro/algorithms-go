package _1872_stone_game_viii

import (
	"fmt"
	"testing"
)

func Test_stoneGameVIII(t *testing.T) {
	tests := []struct {
		stones []int
		want   int
	}{
		{[]int{-1, 2, -3, 4, -5}, 5},
		{[]int{7, -6, 5, 10, 5, -2, -6}, 13},
		{[]int{-10, -12}, -22},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.stones), func(t *testing.T) {
			if got := stoneGameVIII(tt.stones); got != tt.want {
				t.Errorf("stoneGameVIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
