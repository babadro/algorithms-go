package _088_merge_sorted_array

import (
	"github.com/babadro/algorithms-go/slices"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1, 3, 5, 0, 0, 0}, 3, []int{2, 4, 6}, 3, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 1, 0, 0, 0}, 2, []int{2}, 1, []int{1, 1, 2, 0, 0}},
	}

	for i, c := range cases {
		merge(c.nums1, c.m, c.nums2, c.n)
		if !slices.IntSlicesAreEqual(c.nums1, c.expected) {
			t.Errorf("case#%d, want %v, got %v", i+1, c.expected, c.nums1)
		}
	}
}
