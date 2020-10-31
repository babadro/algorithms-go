package _746_min_cost_climbing_stairs

// passed. tptl
func minCostClimbingStairs2(cost []int) int {
	a, b := 0, 0
	for i := len(cost) - 1; i >= 0; i-- {
		res := cost[i] + min(a, b)
		a, b = res, a
	}

	return min(a, b)
}
