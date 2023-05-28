package _0528_random_pick_with_weight

import (
	"math/rand"
	"sort"
)

type Solution struct {
	intervalEnds []int
	num          int
}

// #bnsrg
func Constructor(w []int) Solution {
	num := 0
	intervals := make([]int, len(w))
	for i, weight := range w {
		num += weight
		intervals[i] = num
	}

	return Solution{intervalEnds: intervals, num: num}
}

func (this *Solution) PickIndex() int {
	randomNum := rand.Intn(this.num)

	return sort.Search(len(this.intervalEnds), func(i int) bool {
		return randomNum < this.intervalEnds[i]
	})
}
