package _1722_minimize_hamming_distance_after_swap_operations

import (
	"github.com/babadro/algorithms-go/base/unionFind"
)

// todo 1
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	union := unionFind.NewWQUPC(n)
	for _, pair := range allowedSwaps {
		union.Union(pair[0], pair[1])
	}

	sets := make(map[int][]int)
	for idx := range source {
		root := union.Root(idx)
		sets[root] = append(sets[root], idx)
	}

	counter := 0
	for _, set := range sets {
		mSource, mTarget := make(map[int]bool), make(map[int]bool)
	}

	return counter
}
