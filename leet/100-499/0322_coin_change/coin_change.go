package _322_coin_change

import "math"

// best solution (bottom to up approach). tptl. passed
// todo 2 learn by heart
func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = max
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func coinChangeBruteForce(coins []int, amount int) int {
	return bruteForceHelper(0, coins, amount)
}

func bruteForceHelper(idxCoin int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if idxCoin < len(coins) {
		maxVal := amount / coins[idxCoin]
		minCost := math.MaxInt64
		for x := 0; x <= maxVal; x++ {
			if amount >= x*coins[idxCoin] {
				res := bruteForceHelper(idxCoin+1, coins, amount-x*coins[idxCoin])
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

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
