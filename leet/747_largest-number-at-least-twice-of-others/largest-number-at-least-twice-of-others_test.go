package _747_largest_number_at_least_twice_of_others

import (
	"fmt"
	"testing"
)

func Test_dominantIndex(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 6, 1, 0}, 1},
		{[]int{1, 2, 3, 4}, -1},
		{[]int{1, 0}, 0},
		{[]int{2, 4, 1}, 1},
		{[]int{1}, 0},
		{[]int{2}, -1},
		{[]int{0, 3, 1, 2}, -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := dominantIndex(tt.nums); got != tt.want {
				t.Errorf("dominantIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
