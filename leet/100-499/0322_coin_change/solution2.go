package _322_coin_change

import "math"

func coinChange2(coins []int, amount int) int {
	return helper(0, coins, amount)
}

func helper(idxCoin int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if idxCoin < len(coins) {
		maxVal := amount / coins[idxCoin]
		minCost := math.MaxInt64

		for x := 0; x <= maxVal; x++ {
			if amount >= x*coins[idxCoin] {
				res := helper(idxCoin+1, coins, amount-x*coins[idxCoin])
				if res != -1 {
					minCost = min(minCost, res+x)
				}
			}
		}

		if minCost == math.MaxInt64 {
			return -1
		}
		return minCost
	}

	return -1
}
