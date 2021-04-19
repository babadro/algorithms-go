package _950_reveal_cards_in_increasing_order

import (
	"sort"
)

// passed. medium. not mine. tptl
func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	deque, res := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		deque[i] = i
	}

	sort.Ints(deck)

	for _, card := range deck {
		front := deque[0]
		deque = deque[1:]
		res[front] = card

		if len(deque) > 0 {
			front = deque[0]
			deque = deque[1:]
			deque = append(deque, front)
		}
	}

	return res
}
