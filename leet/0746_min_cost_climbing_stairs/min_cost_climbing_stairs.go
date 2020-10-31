package _746_min_cost_climbing_stairs

// time limit exceed
func minCostClimbingStairs(cost []int) int {
	return min(helper(0, cost), helper(1, cost))
}

func helper(i int, cost []int) int {
	if i == len(cost)-1 || i == len(cost)-2 {
		return cost[i]
	}

	return cost[i] + min(helper(i+1, cost), helper(i+2, cost))
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
