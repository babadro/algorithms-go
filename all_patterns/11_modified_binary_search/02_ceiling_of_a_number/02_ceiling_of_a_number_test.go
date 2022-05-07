package _2_ceiling_of_a_number

import (
	"fmt"
	"testing"
)

func Test_searchCeilingOfANumber(t *testing.T) {
	tests := []struct {
		arr  []int
		key  int
		want int
	}{
		{[]int{4, 6, 10}, 6, 1},
		{[]int{1, 3, 8, 10, 15}, 12, 4},
		{[]int{4, 6, 10}, 17, -1},
		{[]int{4, 6, 10}, -1, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			if got := searchCeilingOfANumber(tt.arr, tt.key); got != tt.want {
				t.Errorf("searchCeilingOfANumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
