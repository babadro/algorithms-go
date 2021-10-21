package _322_coin_change

import (
	"fmt"
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

	if res := recTopDown(dp, coins, 0, 0, amount, nil); res != math.MaxInt64 {
		return res
	}

	return -1
}

func recTopDown(dp [][]int, coins []int, idx, cur, amount int, arr []int) int {
	if amount == 0 {
		return cur
	}

	if idx == len(coins) || amount < 0 {
		return math.MaxInt64
	}

	if idx == 2 && amount == 5 {
		fmt.Println(arr)
	}

	if dp[idx][amount] == -1 {
		sum1 := recTopDown(dp, coins, idx, cur+1, amount-coins[idx], append(arr, coins[idx]))
		sum2 := recTopDown(dp, coins, idx+1, cur, amount, arr)

		if idx == 2 && amount == 5 {
			fmt.Println("ho")
		}

		dp[idx][amount] = min(sum1, sum2)
	}

	return dp[idx][amount]
}

func check(arr []int) bool {
	if len(arr) != 3 {
		return false
	}
	return arr[0] == 1 && arr[1] == 5 && arr[2] == 5
}
