package _198_house_robber

import (
	"fmt"
	"testing"
)

func Test_rob2(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 5, 1, 3, 6, 2, 4}, 15},
		{[]int{2, 10, 14, 8, 1}, 18},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := robTopDown(tt.nums); got != tt.want {
				t.Errorf("robBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
