package _0435_non_overlapping_intervals

import (
	"sort"
)

var max = 0

// todo 1
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)

	helper(0, 0, 0, nil, intervals, &res)

	return len(intervals) - max
}

func helper(lastFinish, i, count int, curr, intervals [][]int, res *[][]int) {
	if i == len(intervals) {
		tmp := make([][]int, len(curr))
		copy(tmp, curr)
		*res = append(*res, tmp...)

		if count > max {
			max = count
		}

		return
	}

	helper(lastFinish, i+1, count, curr, intervals, res)
	if start := intervals[i][0]; start >= lastFinish {
		count++
		lastFinish = intervals[i][1]

		curr = append(curr, intervals[i])

		helper(lastFinish, i+1, count, curr, intervals, res)
	}
}
