package _056_merge_intervals

import "sort"

// tptl passed
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		next, prev := intervals[i], res[len(res)-1]

		if next[0] > prev[1] {
			res = append(res, next)
		} else if next[1] > prev[1] {
			res[len(res)-1][1] = next[1]
		}
	}

	return res
}
