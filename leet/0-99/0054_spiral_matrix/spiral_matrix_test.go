package _054_spiral_matrix

import (
	"testing"

	"github.com/babadro/algorithms-go/matrix"
	"github.com/babadro/algorithms-go/slices"
)

func TestSpiralOrder(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected []int
	}{
		//{matrix.New(0, 1), []int{}},
		//{matrix.New(0, 2), []int{}},
		//{matrix.New(0, 0), []int{}},
		{matrix.New(1, 1), []int{1}},
		{matrix.New(2, 2), []int{1, 2, 4, 3}},
		{matrix.New(3, 3), []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{matrix.New(4, 4), []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}},
		{matrix.New(4, 1), []int{1, 2, 3, 4}},
		{matrix.New(1, 4), []int{1, 2, 3, 4}},
		{matrix.New(2, 3), []int{1, 2, 4, 6, 5, 3}},
		{matrix.New(3, 2), []int{1, 2, 3, 6, 5, 4}},
		{matrix.New(3, 4), []int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8}},
		{matrix.New(4, 3), []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for i, c := range cases {
		if fact := spiralOrder2(c.input); !slices.IntSlicesAreEqual(c.expected, fact) {
			t.Errorf("case#%d, want %v, got %v", i+1, c.expected, fact)
		}
	}
}
