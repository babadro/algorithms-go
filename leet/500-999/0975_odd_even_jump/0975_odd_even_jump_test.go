package _975_odd_even_jump

import (
	"fmt"
	"testing"
)

func Test_oddEvenJumps(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{10, 13, 12, 14, 15}, 2},
		{[]int{2, 3, 1, 1, 4}, 3},
		{[]int{5, 1, 3, 4, 2}, 3},
		{[]int{81, 54, 96, 60, 58}, 2},
		{tleInput, 15366},
	}
	for _, tt := range tests {
		limit := len(tt.arr)
		if limit > 10 {
			limit = 10
		}

		t.Run(fmt.Sprintf("%v", tt.arr[:limit]), func(t *testing.T) {
			if got := oddEvenJumps(tt.arr); got != tt.want {
				t.Errorf("oddEvenJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
