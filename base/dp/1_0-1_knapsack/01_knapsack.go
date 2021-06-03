package __0_1_knapsack

import "github.com/babadro/algorithms-go/utils"

func KnapsackBruteforce(weights, profits []int, capacity int) int {
	return knapsackRecursive(profits, weights, capacity, 0)
}

func knapsackRecursive(profits, weights []int, capacity, currentIndex int) int {
	if capacity <= 0 || currentIndex >= len(profits) {
		return 0
	}

	profit1 := 0
	if weights[currentIndex] <= capacity {
		profit1 = profits[currentIndex] + knapsackRecursive(profits, weights, capacity-weights[currentIndex], currentIndex+1)
	}

	profit2 := knapsackRecursive(profits, weights, capacity, currentIndex+1)

	return utils.Max(profit1, profit2)
}

func KnapsackTopDown(weights, profits []int, capacity int) int {
	dp := make([][]int, len(profits))
	for i := range dp {
		dp[i] = make([]int, capacity+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return topDownRec(dp, profits, weights, capacity, 0)
}

func topDownRec(dp [][]int, profits, weights []int, capacity, currentIndex int) int {
	if capacity <= 0 || currentIndex >= len(profits) {
		return 0
	}

	if dp[currentIndex][capacity] != -1 {
		return dp[currentIndex][capacity]
	}

	profit1 := 0
	if weights[currentIndex] <= capacity {
		profit1 = profits[currentIndex] + topDownRec(dp, profits, weights, capacity-weights[currentIndex], currentIndex+1)
	}

	profit2 := topDownRec(dp, profits, weights, capacity, currentIndex+1)

	dp[currentIndex][capacity] = utils.Max(profit1, profit2)

	return dp[currentIndex][capacity]
}

func KnapsackBottomUp(weights, profits []int, capacity int) int {
	n := len(weights)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for c := 0; c <= capacity; c++ {
		if weights[0] <= c {
			dp[0][c] = profits[0]
		}
	}

	for i := 1; i < n; i++ {
		for c := 1; c <= capacity; c++ {
			profit1 := dp[i-1][c]
			profit2 := 0

			if weights[i] <= c {
				profit2 = profits[i] + dp[i-1][c-weights[i]]
			}

			dp[i][c] = utils.Max(profit1, profit2)
		}
	}

	return dp[n-1][capacity]
}
