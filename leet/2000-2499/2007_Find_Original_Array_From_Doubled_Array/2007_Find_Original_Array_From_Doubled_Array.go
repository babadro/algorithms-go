package _2007_Find_Original_Array_From_Doubled_Array

import (
	"math"
	"sort"
)

// bnsrg todo doesn't work
func findOriginalArray(changed []int) []int {
	sort.Ints(changed)

	m := make(map[int]int)
	for i, num := range changed {
		if _, ok := m[num]; !ok {
			m[num] = i
		}
	}

	var res []int
	for i, num := range changed {
		if num == math.MaxInt64 {
			continue
		}

		twice := num * 2
		idx, ok := m[twice]
		if !ok || changed[idx] != twice {
			return nil
		}

		changed[idx], changed[i] = math.MaxInt64, math.MaxInt64
		m[twice]++

		res = append(res, num)
	}

	return res
}
