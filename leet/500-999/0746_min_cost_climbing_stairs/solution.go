package _746_min_cost_climbing_stairs

// passed. tptl. best solution
func minCostClimbingStairs2(cost []int) int {
	a, b := 0, 0
	for i := len(cost) - 1; i >= 0; i-- {
		res := cost[i] + min(a, b)
		a, b = res, a
	}

	return min(a, b)
}

// TLE, but correct
func minCostClimbingStairsBruteForce(cost []int) int {
	return min(recBruteForce(cost, 0), recBruteForce(cost, 1))
}

func recBruteForce(cost []int, idx int) int {
	if idx >= len(cost) {
		return 0
	}

	return cost[idx] + min(recBruteForce(cost, idx+1), recBruteForce(cost, idx+2))
}

// passed
func minCostClimbingStairsTopDown(cost []int) int {
	dp := make([]int, len(cost))
	for i := range dp {
		dp[i] = -1
	}

	return min(recTopDown(dp, cost, 0), recTopDown(dp, cost, 1))
}

func recTopDown(dp, cost []int, idx int) int {
	if idx >= len(cost) {
		return 0
	}

	if dp[idx] == -1 {
		dp[idx] = cost[idx] + min(recTopDown(dp, cost, idx+1), recTopDown(dp, cost, idx+2))
	}

	return dp[idx]
}

func minCostClimbingStairsBottomUp(cost []int) int {
	dp := make([]int, len(cost)+1)

	dp[0], dp[1] = 0, 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}
