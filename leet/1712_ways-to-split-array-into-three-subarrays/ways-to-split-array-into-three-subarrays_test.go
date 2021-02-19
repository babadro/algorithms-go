package _1712_ways_to_split_array_into_three_subarrays

import (
	"fmt"
	"testing"
)

func Test_waysToSplit(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		//{[]int{1,1,1}, 1},
		//{[]int{1,2,2,2,5,0}, 3},
		//{[]int{3,2,1}, 0},
		//{[]int{2, 5, 0}, 0},
		{[]int{0, 0, 0}, 1}, // todo 1 fails.
		//{[]int{0, 0, 0, 0, 0, 0}, some num }, probably will failed
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := waysToSplit(tt.nums); got != tt.want {
				t.Errorf("waysToSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
