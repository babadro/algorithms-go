package _056_merge_intervals

import (
	"sort"

	"github.com/babadro/algorithms-go/utils"
)

func merge2(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}

		return false
	})

	for i := 1; i < len(intervals); i++ {
		left, right := intervals[i-1], intervals[i]
		if res, ok := overlap(left, right); ok {
			intervals[i] = res
			intervals[i-1][0] = -1
		}
	}

	idx := 0
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] != -1 {
			intervals[idx] = intervals[i]
			idx++
		}
	}

	return intervals[:idx]
}

func overlap(i1, i2 []int) ([]int, bool) {
	if i1[1] >= i2[0] {
		return []int{i1[0], utils.Max(i1[1], i2[1])}, true
	}

	return []int{}, false
}
