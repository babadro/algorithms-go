package _015_3Sum

import (
	"github.com/babadro/algorithms-go/slices"
	"testing"
)

func TestThreeeSum(t *testing.T) {
	cases := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{}, nil},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, 0, 1}, {-1, -1, 2}}},
	}

	for i, c := range cases {
		fact := threeSum(c.nums)
		if !slices.SliceOfIntSlicesAreEqual(fact, c.expected) {
			t.Errorf("case#%d, want %v, got %v", i+1, c.expected, fact)
		}
	}
}
