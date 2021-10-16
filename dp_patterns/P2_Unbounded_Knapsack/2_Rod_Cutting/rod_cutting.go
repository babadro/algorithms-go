package __Rod_Cutting

func RodCuttingBruteForce(lengths, prices []int, n int) int {
	return recBruteForce(lengths, prices, 0, 0, n)
}

func recBruteForce(lengths, prices []int, profit, idx, n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 || idx == len(lengths) {
		return profit
	}

	res1 := profit + recBruteForce(lengths, prices, prices[idx], idx, n-lengths[idx])
	res2 := profit + recBruteForce(lengths, prices, 0, idx+1, n)

	return max(res1, res2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func RodCuttingBottomUp(lengths, prices []int, n int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(prices); i++ {
		for length := 1; length <= n; length++ {
			var res1, res2 int
			if i > 0 {
				res1 = dp[i-1][length]
			}

			if length >= lengths[i] {
				res2 = prices[i] + dp[i][length-lengths[i]]
			}

			dp[i][length] = max(res1, res2)
		}
	}

	return dp[len(prices)-1][n]
}
