package _283_move_zeroes

import (
	"github.com/babadro/algorithms-go/slices"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{0}, []int{0}},
		{[]int{1, 0}, []int{1, 0}},
		{[]int{0, 1}, []int{1, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{3, 0, 0, 1, 4, 5, 0, 1}, []int{3, 1, 4, 5, 1, 0, 0, 0}},
	}

	for i, c := range cases {
		moveZeroes(c.input)
		if !slices.IntSlicesAreEqual(c.input, c.expected) {
			t.Errorf("case#%d, want %v, got %v", i+1, c.expected, c.input)
		}
	}
}
