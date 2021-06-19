package _322_coin_change

import "sort"

// doesn't work, just for illustrate greedy approach.
func coinChangeGreedy(coins []int, amount int) int {
	sort.Ints(coins)

	res := 0
	for i := len(coins) - 1; i >= 0; i-- {
		for amount >= coins[i] {
			amount -= coins[i]
			res++
		}
	}

	if amount == 0 {
		return res
	}

	return -1
}
