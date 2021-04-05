package _35

import (
	"fmt"
	"testing"
)

func Test_countDifferentSubsequenceGCDs(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{6, 10, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := countDifferentSubsequenceGCDs(tt.nums); got != tt.want {
				t.Errorf("countDifferentSubsequenceGCDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
