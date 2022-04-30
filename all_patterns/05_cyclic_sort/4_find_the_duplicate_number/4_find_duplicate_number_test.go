package __find_the_duplicate_number

import (
	"fmt"
	"testing"
)

func Test_findNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		//{[]int{1, 4, 4, 3, 2}, 4},
		//{[]int{2, 1, 3, 3, 5, 4}, 3},
		//{[]int{2, 4, 1, 4, 4}, 4},
		{[]int{3, 2, 3, 1, 4}, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := findNumberTwoPointers(tt.nums); got != tt.want {
				t.Errorf("findNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
