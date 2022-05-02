package __Introduction

import (
	"fmt"
	"testing"
)

func Test_findMissingNumber(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 5, 2, 6, 4}, 3},
		{[]int{2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 5, 6, 7}, 4},
		{[]int{}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			if got := findMissingNumber(tt.arr); got != tt.want {
				t.Errorf("findMissingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
