package _084_largest_rectangle_in_histogram

import (
	"fmt"
	"testing"
)

func Test_largestRectangleArea(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.heights), func(t *testing.T) {
			if got := largestRectangleArea(tt.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
