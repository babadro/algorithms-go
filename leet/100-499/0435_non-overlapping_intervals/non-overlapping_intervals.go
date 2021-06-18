package _0435_non_overlapping_intervals

import "math"

var min = math.MaxInt64

// todo 1
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)

	return n

}

func helper2(lastFinish, i, count int, intervals [][]int) {
	if i == len(intervals) {
		if count < min {
			min = count
		}

		return
	}

	//if lastFinish < intervals[i][0]
}
