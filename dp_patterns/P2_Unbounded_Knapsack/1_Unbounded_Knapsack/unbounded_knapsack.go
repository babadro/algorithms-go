package __Unbounded_Knapsack

func knapsackBruteForce(profits, weights []int, capacity int) int {
	return recBruteForce(profits, weights, capacity, 0)
}

func recBruteForce(profits, weights []int, capacity, idx int) int {
	if capacity <= 0 || idx == len(profits) {
		return 0
	}

	var profit1, profit2 int
	if weights[idx] <= capacity {
		profit1 = profits[idx] + recBruteForce(profits, weights, capacity-weights[idx], idx)
	}

	profit2 = recBruteForce(profits, weights, capacity, idx+1)

	return max(profit1, profit2)
}

func knapsackTopDown(profits, weights []int, capacity int) int {
	dp := make([][]int, len(profits))
	for i := range dp {
		dp[i] = make([]int, capacity+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(dp, profits, weights, capacity, 0)
}

func recTopDown(dp [][]int, profits, weights []int, capacity, idx int) int {
	if capacity <= 0 || idx == len(profits) {
		return 0
	}

	if dp[idx][capacity] == -1 {
		var profit1, profit2 int
		if weights[idx] <= capacity {
			profit1 = profits[idx] + recTopDown(dp, profits, weights, capacity-weights[idx], idx)
		}

		profit2 = recTopDown(dp, profits, weights, capacity, idx+1)

		dp[idx][capacity] = max(profit1, profit2)
	}

	return dp[idx][capacity]
}

func knapsackBottomUp(profits, weights []int, capacity int) int {
	n := len(profits)
	if capacity == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 0; i < n; i++ {
		for c := 1; c <= capacity; c++ {
			var profit1, profit2 int
			if i > 0 {
				profit1 = dp[i-1][c]
			}

			if weights[i] <= c {
				profit2 = profits[i] + dp[i][c-weights[i]]
			}

			dp[i][c] = max(profit1, profit2)
		}
	}

	return dp[n-1][capacity]
}

// KnapsackBottomUpSpaceOptimized solution is also possible. Look KnapsackBottomUpSpaceOptimized
// in Pattern_1 0_1_Knapsack

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// todo 2 doesn't work
func FindSelectedElements(dp [][]int, weights, profits []int, capacity int) []int {
	var res []int
	n := len(profits)
	totalProfit := dp[n-1][capacity]
	for i := n - 1; i >= 0; i-- {
		for capacity != 0 {
			if totalProfit != dp[i-1][capacity] {
				res = append(res, i)
				capacity -= weights[i]
				totalProfit -= profits[i]
			}
		}
	}

	if totalProfit != 0 {
		res = append(res, 0)
	}

	return res
}
