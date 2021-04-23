package _679_24_game

import (
	"fmt"
	"testing"
)

func Test_judgePoint24(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{4, 1, 8, 7}, true},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%d%d%d%d", tt.nums[0], tt.nums[1], tt.nums[2], tt.nums[3])
		t.Run(name, func(t *testing.T) {
			if got := judgePoint24(tt.nums); got != tt.want {
				t.Errorf("judgePoint24() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genPerm(t *testing.T) {
	tests := [][]int{
		{1, 2, 3, 4},
		{1},
		{1, 2},
		{1, 2, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			got := genPerm(tt)
			for _, perm := range got {
				t.Log(perm)
			}
		})
	}
}
