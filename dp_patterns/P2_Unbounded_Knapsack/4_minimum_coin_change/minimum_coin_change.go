// look leetcode 322 for optimized memory solution
package __minimum_coin_change

import (
	"fmt"
	"math"
)

func coinChangeBruteForce(coins []int, amount int) int {
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

// passed.
func coinChangeTopDown(coins []int, amount int) int {
	dp := make([][]int, len(coins))
	for i := range dp {
		dp[i] = make([]int, amount+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	if res := recTopDown(dp, coins, 0, amount); res != math.MaxInt64 {
		return res
	}

	return -1
}

func recTopDown(dp [][]int, coins []int, idx, amount int) int {
	if amount == 0 {
		return 0
	}

	if idx == len(coins) {
		return math.MaxInt64
	}

	if dp[idx][amount] == -1 {
		count1 := math.MaxInt64
		if coins[idx] <= amount {
			res := recTopDown(dp, coins, idx, amount-coins[idx])
			if res != math.MaxInt64 {
				count1 = res + 1
			}
		}

		count2 := recTopDown(dp, coins, idx+1, amount)

		dp[idx][amount] = min(count1, count2)
	}

	return dp[idx][amount]
}

func coinChangeBottomUp(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64
		}
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 0
	}

	for i := 0; i < n; i++ {
		for a := 0; a <= amount; a++ {
			if i > 0 {
				dp[i][a] = dp[i-1][a]
			}

			if a >= coins[i] {
				if dp[i][a-coins[i]] != math.MaxInt64 {
					dp[i][a] = min(dp[i][a], dp[i][a-coins[i]]+1)
				}
			}
		}
	}

	if res := dp[n-1][amount]; res != math.MaxInt64 {
		return res
	}

	return -1
}

// todo 1 doesn't work. Just for demonstration
func coinChangeTopDown3(coins []int, amount int) int {
	dp := make([][]int, len(coins))
	for i := range dp {
		dp[i] = make([]int, amount+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	if res := recTopDown3(dp, coins, 0, 0, 0, amount); res != math.MaxInt64 {
		return res
	}

	return -1
}

func recTopDown3(dp [][]int, coins []int, idx, curCount, curAmount, target int) int {
	if curAmount == target {
		return curCount
	}

	if idx == len(coins) || curAmount > target {
		return math.MaxInt64
	}

	if dp[idx][curAmount] == -1 {
		// don't pass curCount to func. This approach doesn't work
		sum1 := recTopDown3(dp, coins, idx, curCount+1, curAmount+coins[idx], target)
		sum2 := recTopDown3(dp, coins, idx+1, curCount, curAmount, target)

		dp[idx][target] = min(sum1, sum2)
	}

	return dp[idx][target]
}

type item struct {
	count int
	arr   []int
}

// it works. arr in dp just for demonstration and debug.
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

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
