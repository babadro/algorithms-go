package __Khapsack

import "github.com/babadro/algorithms-go/utils"

func Knapsack(profits, weights []int, capacity int) int {
	dp := make([][]int, len(profits))
	for i := range dp {
		dp[i] = make([]int, capacity+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	res := req(dp, profits, weights, capacity, 0)
	return res
}

func req(dp [][]int, profits, weights []int, capacity, currentIndex int) int {
	if capacity <= 0 || currentIndex >= len(profits) {
		return 0
	}

	if dp[currentIndex][capacity] != -1 {
		return dp[currentIndex][capacity]
	}

	profit1 := 0
	if weights[currentIndex] <= capacity {
		profit1 = profits[currentIndex] + req(dp, profits, weights, capacity-weights[currentIndex], currentIndex+1)
	}

	profit2 := req(dp, profits, weights, capacity, currentIndex+1)
	dp[currentIndex][capacity] = utils.Max(profit1, profit2)

	return dp[currentIndex][capacity]
}
