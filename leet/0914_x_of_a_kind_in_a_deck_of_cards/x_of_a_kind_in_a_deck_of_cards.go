package _914_x_of_a_kind_in_a_deck_of_cards

import "sort"

// todo 1
func hasGroupsSizeX(deck []int) bool {
	n := len(deck)
	if n%2 != 0 {
		return false
	}

	sort.Ints(deck)

	minLen, currLen := 0, 1

	for i := 1; i <= n; i++ {
		if i < n && deck[i] == deck[i-1] {
			currLen++
			continue
		}

		if currLen == 1 {
			return false
		}

		if minLen == 0 || currLen < minLen {
			if minLen%currLen != 0 {
				return false
			}

			minLen = currLen

			continue
		}

		if currLen%minLen != 0 {
			return false
		}

		currLen = 0
	}

	return true
}
