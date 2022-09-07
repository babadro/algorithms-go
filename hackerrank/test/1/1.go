package _1

import (
	"sort"
)

func bioHazard(n int32, allergic []int32, poisonous []int32) int64 {
	poisDict := make(map[int32][]int32)
	for i := range allergic {
		a, b := swap(allergic[i], poisonous[i])

		poisDict[b] = append(poisDict[b], a)
	}

	for _, v := range poisDict {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}

	var res int64
	l := int32(1)
	for r := int32(1); r <= n; r++ {
		if len(poisDict[r]) == 0 {
			res += int64(r - l + 1)
			continue
		}

		lastValid := poisDict[r][len(poisDict[r])-1] + 1
		if lastValid > l {
			l = lastValid
		}

		res += int64(r - l + 1)
	}

	return res
}

func swap(a, b int32) (int32, int32) {
	if a > b {
		return b, a
	}

	return a, b
}
