package _350_intersection_of_two_arrays_2

import (
	"github.com/babadro/algorithms-go/slices"
	"sort"
	"testing"
)

func TestIntersect(t *testing.T) {
	cases := []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{1}, []int{}},
		{[]int{1}, []int{}, []int{}},
		{[]int{1, 2}, []int{}, []int{}},
		{[]int{1, 2}, []int{3, 4}, []int{}},
		{[]int{1, 2, 3, 4}, []int{3, 4}, []int{3, 4}},
		{[]int{1, 2, 3, 4}, []int{4, 5, 6}, []int{4}},
		{[]int{1, 2, 2, 2}, []int{1, 2, 2, 5}, []int{1, 2, 2}},
	}

	for i, c := range cases {
		fact := intersect(c.nums1, c.nums2)
		sort.Ints(c.expected)
		sort.Ints(fact)
		if !slices.IntSlicesAreEqual(fact, c.expected) {
			t.Errorf("case#%d, want %v, got %v", i+1, c.expected, fact)
		}
	}
}
