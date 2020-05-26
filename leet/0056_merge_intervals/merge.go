package _056_merge_intervals

import "sort"

// 8 ms 96%, 4.7MB 100 %
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := len(res) - 1
		if res[last][1] >= intervals[i][0] {
			if res[last][1] < intervals[i][1] {
				res[last][1] = intervals[i][1]
			}
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
