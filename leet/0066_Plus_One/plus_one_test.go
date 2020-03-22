package _066_Plus_One

import (
	"github.com/babadro/algorithms-go/slices"
	"testing"
)

func TestPlusOne(t *testing.T) {
	cases := []struct {
		digits   []int
		expected []int
	}{
		{[]int{0}, []int{1}},
		{[]int{1, 2}, []int{1, 3}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{1, 0, 9, 9}, []int{1, 1, 0, 0}},
	}

	for i, c := range cases {
		if fact := plusOne(c.digits); !slices.IntSlicesAreEqual(fact, c.expected) {
			t.Errorf("case#%d, want %v, got %v", i+1, c.expected, fact)
		}
	}
}
