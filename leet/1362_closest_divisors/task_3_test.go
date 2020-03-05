package _1362_closest_divisors

import (
	"github.com/babadro/algorithms-go/slices"
	"testing"
)

func TestDivisors(t *testing.T) {
	cases := []struct {
		input    int
		expected []int
	}{
		{8, []int{3, 3}},
		{123, []int{5, 25}},
		{999, []int{40, 25}},
	}

	for i, c := range cases {
		fact := closestDivisors(c.input)
		if !slices.IntSlicesAreEqual(fact, c.expected) {
			swap(fact)
			if !slices.IntSlicesAreEqual(fact, c.expected) {
				t.Errorf("case#%d, want %v, got %v", i+1, c.expected, fact)
			}
		}
	}
}

func swap(arr []int) {
	arr[0], arr[1] = arr[1], arr[0]
}
