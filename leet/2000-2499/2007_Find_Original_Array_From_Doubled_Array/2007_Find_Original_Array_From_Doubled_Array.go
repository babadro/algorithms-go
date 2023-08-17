package _2007_Find_Original_Array_From_Doubled_Array

import (
	"math"
	"sort"
)

// bnsrg
// passed, but not so fast. todo 2 find faster solution, probably without sorting
func findOriginalArray(changed []int) []int {
	if len(changed)%2 != 0 {
		return nil
	}

	sort.Ints(changed)

	m := make(map[int]int)
	for i, num := range changed {
		if _, ok := m[num]; !ok {
			m[num] = i
		}
	}

	var res []int
	for _, num := range changed {
		if num == math.MaxInt64 {
			continue
		}

		if num == 0 {
			m[0]++
		}

		twice := num * 2
		idx, ok := m[twice]
		if !ok || idx >= len(changed) || changed[idx] != twice {
			return nil
		}

		changed[idx] = math.MaxInt64
		m[twice]++

		res = append(res, num)
	}

	return res
}
