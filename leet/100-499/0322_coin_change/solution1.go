package _322_coin_change

import (
	"math"
)

func coinChangeBruteForce2(coins []int, amount int) int {
	if res := recBruteForce(coins, 0, 0, amount); res != math.MaxInt64 {
		return res
	}

	return -1
}

func recBruteForce(coins []int, idx, cur, amount int) int {
	if amount == 0 {
		return cur
	}

	if idx == len(coins) || amount < 0 {
		return math.MaxInt64
	}

	sum1 := recBruteForce(coins, idx, cur+1, amount-coins[idx])
	sum2 := recBruteForce(coins, idx+1, cur, amount)

	return min(sum1, sum2)
}

// todo 1 doesn't work
func coinChangeTopDown(coins []int, amount int) int {
	dp := make([][]int, len(coins))
	for i := range dp {
		dp[i] = make([]int, amount+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	if res := recTopDown(dp, coins, 0, 0, amount); res != math.MaxInt64 {
		return res
	}

	return -1
}

func recTopDown(dp [][]int, coins []int, idx, cur, amount int) int {
	if amount == 0 {
		return cur
	}

	if idx == len(coins) || amount < 0 {
		return math.MaxInt64
	}

	if dp[idx][cur] == -1 {
		sum1 := recTopDown(dp, coins, idx, cur+1, amount-coins[idx])
		sum2 := recTopDown(dp, coins, idx+1, cur, amount)

		dp[idx][cur] = min(sum1, sum2)
	}

	return dp[idx][cur]
}
