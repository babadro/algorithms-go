package __house_thief

import (
	"fmt"
	"testing"
)

var a = []int{4, 1, 5, 61, 3, 1, 6, 8, 4, 7, 34, 4, 6, 6, 3, 7, 7, 8, 5, 6, 4, 1, 5, 61, 3, 1, 6, 8, 4, 7, 34, 4, 6, 6, 3, 7, 7, 8, 5, 6, 4, 1, 5, 61, 3, 1}

func Test_findMaxSteal(t *testing.T) {
	tests := []struct {
		wealth []int
		want   int
	}{
		{[]int{2, 5, 1, 3, 6, 2, 4}, 15},
		{[]int{2, 10, 14, 8, 1}, 18},
		{a, 332},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.wealth), func(t *testing.T) {
			if got := findMaxStealBottomUpOptimized(tt.wealth); got != tt.want {
				t.Errorf("findMaxSteal() = %v, want %v", got, tt.want)
			}
		})
	}
}
