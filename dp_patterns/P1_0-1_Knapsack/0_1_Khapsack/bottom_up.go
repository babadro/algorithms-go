package __1_Khapsack

import "github.com/babadro/algorithms-go/utils"

func KnapsackBottomUp(profits, weights []int, capacity int) int {
	n := len(profits)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// fill top row with profits[0] value
	for c := 0; c <= capacity; c++ {
		if weights[0] <= c {
			dp[0][c] = profits[0]
		}
	}

	for i := 1; i < n; i++ {
		for c := 1; c <= capacity; c++ {
			profit1, profit2 := 0, 0
			if weights[i] <= c {
				profit1 = profits[i] + dp[i-1][c-weights[i]]
			}

			profit2 = dp[i-1][c]

			dp[i][c] = utils.Max(profit1, profit2)
		}
	}

	return dp[n-1][capacity]
}

// Here we use two rows only. Space = O(Capacity)
func KnapsackBottomUpSpaceOptimized(profits, weights []int, capacity int) int {
	n := len(profits)
	dp := [2][]int{
		make([]int, capacity+1),
		make([]int, capacity+1),
	}

	// fill top row with profits[0] value
	for c := 0; c <= capacity; c++ {
		if weights[0] <= c {
			dp[0][c] = profits[0]
		}
	}

	for i := 1; i < n; i++ {
		for c := 1; c <= capacity; c++ {
			profit1, profit2 := 0, 0
			if weights[i] <= c {
				profit1 = profits[i] + dp[1][c-weights[i]]
			}

			profit2 = dp[0][c]

			dp[1][c] = utils.Max(profit1, profit2)

			dp[0], dp[1] = dp[1], dp[0] // swap the rows
		}
	}

	return dp[0][capacity]
}
