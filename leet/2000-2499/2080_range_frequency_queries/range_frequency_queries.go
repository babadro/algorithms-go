package _2080_range_frequency_queries

import "sort"

// passed. easy
type RangeFreqQuery struct {
	m map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	m := make(map[int][]int)
	for i, num := range arr {
		m[num] = append(m[num], i)
	}

	return RangeFreqQuery{m}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	idxes := this.m[value]
	start := sort.Search(len(idxes), func(i int) bool {
		return idxes[i] >= left
	})

	finish := sort.Search(len(idxes), func(i int) bool {
		return idxes[i] > right
	})

	return finish - start
}
