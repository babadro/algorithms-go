package _1_counting_money

import "sort"

// leetcode 322
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
