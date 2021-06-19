package _0435_non_overlapping_intervals

import "sort"

// tptl. passed. best solution. medium. #greedy
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	j, counter := 0, 1
	for i := 1; i < len(intervals); i++ {
		if intervals[j][1] <= intervals[i][0] {
			counter++
			j = i
		}
	}

	return len(intervals) - counter
}
