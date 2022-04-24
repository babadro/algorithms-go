package __find_the_missing_number

import (
	"fmt"
	"testing"
)

func Test_findMissingNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{4, 0, 3, 1}, 2},
		{[]int{8, 3, 5, 2, 4, 6, 0, 1}, 7},
		{[]int{0, 1, 2}, 3},
		{[]int{1, 2}, 0},
		{[]int{0, 2}, 1},
		{[]int{4, 3, 2, 1}, 0},
		{[]int{0, 1, 3, 4}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := findMissingNumber(tt.nums); got != tt.want {
				t.Errorf("findMissingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
