package _1848_minimum_distance_to_the_target_element

import (
	"fmt"
	"testing"
)

func Test_getMinDistance(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		start  int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5}, 5, 3, 1},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0, 0},
		{[]int{1}, 1, 0, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := getMinDistance(tt.nums, tt.target, tt.start); got != tt.want {
				t.Errorf("getMinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
