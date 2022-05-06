package _1_merge_intervals

import (
	"sort"

	"github.com/babadro/algorithms-go/utils"
)

// look leetcode 56 - it's a best solution
func merge(intervals [][2]int) [][2]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}

		// len comparison
		return intervals[i][1]-intervals[i][0] > intervals[j][1]-intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		left, right := intervals[i-1], intervals[i]
		if res, ok := overlap(left, right); ok {
			intervals[i] = res
			intervals[i-1][0] = -1
		}
	}

	// there is a mistake
	for i := 0; i < len(intervals); i++ {
		last := len(intervals) - 1
		if intervals[i][0] == -1 {
			intervals[i], intervals[last] = intervals[last], intervals[i]
			intervals = intervals[:last]
		}
	}

	return intervals
}

func overlap(i1, i2 [2]int) ([2]int, bool) {
	if i1[1] > i2[0] {
		return [2]int{i1[0], utils.Max(i1[1], i2[1])}, true
	}

	return [2]int{}, false
}
