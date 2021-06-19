package _0435_non_overlapping_intervals

import (
	"fmt"
	"math"
	"sort"
)

// TLE, but, probably, correct
func eraseOverlapIntervalsTLE(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		start1, start2 := intervals[i][0], intervals[j][0]
		if start1 == start2 {
			return intervals[i][1] < intervals[j][1]
		}

		return start1 < start2
	})

	fmt.Println(intervals)

	var max = 0

	helper(math.MinInt64, 0, 0, intervals, &max)

	return len(intervals) - max
}

func helper(lastFinish, i, count int, intervals [][]int, max *int) {
	if i == len(intervals) {
		if count > *max {
			*max = count
		}

		return
	}

	if len(intervals)-i+count < *max {
		return
	}

	//fmt.Printf("count=%d\n", count)

	if start := intervals[i][0]; start >= lastFinish {
		helper(intervals[i][1], i+1, count+1, intervals, max)
	}

	helper(lastFinish, i+1, count, intervals, max)
}
