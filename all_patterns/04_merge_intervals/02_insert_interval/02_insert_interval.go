package _2_insert_interval

import "github.com/babadro/algorithms-go/utils"

// leetcode 57
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)
	i := 0
	for ; intervals[i][1] < newInterval[0]; i++ {
		res = append(res, intervals[i])
	}

	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = utils.Min(newInterval[0], intervals[i][0])
		newInterval[1] = utils.Max(newInterval[1], intervals[i][1])
		i++
	}

	res = append(res, newInterval)

	for i < len(intervals) {
		res = append(res, intervals[i])
	}

	return res
}
