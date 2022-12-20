package _015_3Sum

import (
	"sort"
	"testing"

	"github.com/babadro/algorithms-go/slices"
	"github.com/stretchr/testify/require"
)

func TestThreeeSum(t *testing.T) {
	cases := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{}, nil},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{[]int{0}, nil},
		{[]int{-1, 0, 1}, [][]int{{-1, 0, 1}}},
		{[]int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
	}

	for _, c := range cases {

		fact := threeSum(c.nums)

		slices.SortEach(fact)
		sort.Slice(fact, func(i, j int) bool {
			return slices.Less(fact[i], fact[j])
		})

		slices.SortEach(c.expected)
		sort.Slice(c.expected, func(i, j int) bool {
			return slices.Less(c.expected[i], c.expected[j])
		})

		require.Equal(t, c.expected, fact)
	}
}
