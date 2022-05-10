package _2_insert_interval

import "github.com/babadro/algorithms-go/utils"

// see leetcode 57
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	res := make([][]int, 0, len(intervals)+1)
	i := 0
	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		res = append(res, intervals[i])
	}

	for ; i < len(intervals) && newInterval[1] >= intervals[i][0]; i++ {
		newInterval[0] = utils.Min(newInterval[0], intervals[i][0])
		newInterval[1] = utils.Max(newInterval[1], intervals[i][1])
	}

	res = append(res, newInterval)

	for ; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}

	return res
}
