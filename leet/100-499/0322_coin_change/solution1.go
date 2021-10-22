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

	if res := recTopDown(dp, coins, 0, 0, 0, amount); res != math.MaxInt64 {
		return res
	}

	return -1
}

func recTopDown(dp [][]int, coins []int, idx, curCount, curAmount, target int) int {
	if curAmount == target {
		return curCount
	}

	if idx == len(coins) || curAmount > target {
		return math.MaxInt64
	}

	if dp[idx][curAmount] == -1 {
		sum1 := recTopDown(dp, coins, idx, curCount+1, curAmount+coins[idx], target)
		sum2 := recTopDown(dp, coins, idx+1, curCount, curAmount, target)

		dp[idx][target] = min(sum1, sum2)
	}

	return dp[idx][target]
}

type item struct {
	count int
	arr   []int
}

func coinChangeTopDown2(coins []int, amount int) int {
	dp := make([][]item, len(coins))
	for i := range dp {
		dp[i] = make([]item, amount+1)
		for j := range dp[i] {
			dp[i][j].count = -1
		}
	}

	if res := recTopDown2(dp, coins, 0, amount, nil); res != math.MaxInt64 {
		return res
	}

	return -1
}

func recTopDown2(dp [][]item, coins []int, idx, amount int, arr []int) int {
	if amount == 0 {
		return 0
	}

	if idx == len(coins) {
		return math.MaxInt64
	}

	if idx == 2 && amount == 5 {
		fmt.Println(arr)
	}

	if dp[idx][amount].count == -1 {
		count1 := math.MaxInt64
		if coins[idx] <= amount {
			res := recTopDown2(dp, coins, idx, amount-coins[idx], append(arr, coins[idx]))
			if res != math.MaxInt64 {
				count1 = res + 1
			}
		}

		count2 := recTopDown2(dp, coins, idx+1, amount, arr)

		if idx == 2 && amount == 5 {
			fmt.Println("ho")
		}

		copyArr := make([]int, len(arr))
		copy(copyArr, arr)

		if count1 < count2 {
			dp[idx][amount].count, dp[idx][amount].arr = count1, copyArr
		} else {
			dp[idx][amount].count, dp[idx][amount].arr = count2, copyArr
		}
	}

	return dp[idx][amount].count
}

func check(arr []int) bool {
	if len(arr) != 3 {
		return false
	}
	return arr[0] == 1 && arr[1] == 5 && arr[2] == 5
}
