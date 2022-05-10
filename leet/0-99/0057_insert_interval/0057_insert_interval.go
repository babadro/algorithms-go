package _0057_insert_interval

// tptl. passed.
// todo 2 reuse intervals for returning result
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
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
	}

	res = append(res, newInterval)

	for ; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
