package _190_rotate_array_go

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
		//{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 3, []int{10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9},},
		{[]int{1, 2, 3, 4}, 3, []int{2, 3, 4, 1}},
	}

	for i, c := range cases {
		rotate(c.nums, c.k)
		if !slices.IntSlicesAreEqual(c.nums, c.expected) {
			t.Errorf("case#%d, want %v, got %v", i+1, c.expected, c.nums)
		}
	}
}
