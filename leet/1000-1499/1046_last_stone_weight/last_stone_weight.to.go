package _1046_last_stone_weight

import "sort"

// heap. tptl. passed. easy.
func lastStoneWeight(stones []int) int {
	sort.Ints(stones)

	for len(stones) > 0 {
		if len(stones) > 1 {
			last, lastButOne := len(stones)-1, len(stones)-2
			newStone := stones[last] - stones[lastButOne]
			stones = stones[:lastButOne]

			if newStone == 0 {
				continue
			}

			idx := sort.SearchInts(stones, newStone)

			// add to sorted
			stones = append(stones, 0)
			copy(stones[idx+1:], stones[idx:len(stones)-1])
			stones[idx] = newStone
		} else {
			return stones[0]
		}
	}

	return 0
}
