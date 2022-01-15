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

type times struct {
	grow, plant []int
}

func (t times) Len() int           { return len(t.plant) }
func (t times) Less(i, j int) bool { return t.plant[i] > t.plant[j] }
func (t times) Swap(i, j int) {
	t.plant[i], t.plant[j] = t.plant[j], t.plant[i]
	t.grow[i], t.grow[j] = t.grow[j], t.grow[i]
}

func earliestFullBloom2(plantTime []int, growTime []int) int {
	t := times{plantTime, growTime}
	sort.Sort(t)

	res, curDay := 0, 0
	for i := range plantTime {
		res = max(res, curDay+plantTime[i]+growTime[i])
		curDay += plantTime[i]
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
