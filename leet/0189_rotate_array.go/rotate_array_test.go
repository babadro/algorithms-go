package _189_rotate_array_go

import (
	"github.com/babadro/algorithms-go/slices"
	"testing"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 2, []int{11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4}, 3, []int{2, 3, 4, 1}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{[]int{1}, 1, []int{1}},
		{[]int{-1}, 2, []int{-1}},
		{[]int{}, 2, []int{}},
		{[]int{1, 2}, 3, []int{2, 1}},
		{[]int{1, 2, 3}, 4, []int{3, 1, 2}},
	}

	for i, c := range cases {
		rotate(c.nums, c.k)
		if !slices.IntSlicesAreEqual(c.nums, c.expected) {
			t.Errorf("case#%d, want %v, got %v", i+1, c.expected, c.nums)
		}
	}
}
