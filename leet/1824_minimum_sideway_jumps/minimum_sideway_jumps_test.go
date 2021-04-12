package _1824_minimum_sideway_jumps

import (
	"fmt"
	"testing"
)

func Test_minSideJumps(t *testing.T) {
	tests := []struct {
		obstacles []int
		want      int
	}{
		{[]int{0, 1, 2, 3, 0}, 2},
		//{[]int{0,1,1,3,3,0}, 0},
		//{[]int{0,2,1,0,3,0}, 2},
		//{[]int{0,0,2,0,0,0,2,1,2,0,0}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.obstacles), func(t *testing.T) {
			if got := minSideJumpsDP(tt.obstacles); got != tt.want {
				t.Errorf("minSideJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
