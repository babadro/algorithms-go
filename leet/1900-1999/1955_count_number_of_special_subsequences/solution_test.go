package _1955_count_number_of_special_subsequences

import (
	"fmt"
	"testing"
)

func Test_countSpecialSubsequences(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{0, 1, 2, 2}, 3},
		{[]int{2, 2, 0, 0}, 0},
		{[]int{0, 1, 2, 0, 1, 2}, 7},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.nums)
		if len(tt.nums) > 10 {
			name = fmt.Sprintf("%v<...>", tt.nums[:10])
		}

		t.Run(name, func(t *testing.T) {
			if got := countSpecialSubsequences(tt.nums); got != tt.want {
				t.Errorf("countSpecialSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
