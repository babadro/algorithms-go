package _01_minimum_operations_to_sort_array

import (
	"fmt"
	"testing"
)

func Test_minOperations(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{5, 1, 4, 3, 2}, 3},
		{[]int{3, 2, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			if got := minOperations(tt.arr); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
