package __fibonacci

// edu
func FibRec(n int) int {
	if n < 2 {
		return n
	}

	return FibRec(n-1) + FibRec(n-2)
}

// edu
func FibTopDown(n int) int {
	memo := make([]int, n+1)

	return topDownRec(memo, n)
}

func topDownRec(memo []int, n int) int {
	if n < 2 {
		return n
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = topDownRec(memo, n-1) + topDownRec(memo, n-2)

	return memo[n]
}

func FibBottomUp(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
