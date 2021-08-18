package __1_Khapsack

func FindSelectedElements(dp [][]int, weights, profits []int, capacity int) []int {
	var res []int
	n := len(profits)
	totalProfit := dp[n-1][capacity]
	for i := n - 1; i > 0; i-- {
		if totalProfit != dp[i-1][capacity] {
			res = append(res, i)
			capacity -= weights[i]
			totalProfit -= profits[i]
		}
	}

	if totalProfit != 0 {
		res = append(res, 0)
	}

	return res
}
