package _2136_earliest_possible_day_full_bloom

import "sort"

// https://leetcode.com/problems/earliest-possible-day-of-full-bloom/discuss/1676837/Grow-then-plant
// tptl passed. hard.
func earliestFullBloom(plantTime []int, growTime []int) int {
	pairs := make([][2]int, len(plantTime))
	for i := range plantTime {
		pairs[i][0], pairs[i][1] = growTime[i], plantTime[i]
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	res := 0
	for _, pair := range pairs {
		growT, plantT := pair[0], pair[1]
		res = max(res, growT) + plantT
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
