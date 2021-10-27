package __number_factors

func CountWaysBruteForce(n int) int {
	if n == 0 {
		return 1 // {}
	}

	if n == 1 {
		return 1 // {1}
	}

	if n == 2 {
		return 1 // {1,1}
	}

	if n == 3 {
		return 2 // {1,1,1}, {3}
	}

	return CountWaysBruteForce(n-1) + CountWaysBruteForce(n-3) + CountWaysBruteForce(n-4)
}

func CountWaysTopDown(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}

	return recTopDown(dp, n)
}

func recTopDown(dp []int, n int) int {
	if n < 3 {
		return 1
	}

	if n == 3 {
		return 2
	}

	if dp[n] == -1 {
		dp[n] = recTopDown(dp, n-1) + recTopDown(dp, n-3) + recTopDown(dp, n-4)
	}

	return dp[n]
}

func CountWaysBottomUp(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2], dp[3] = 1, 1, 1, 2

	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-3] + dp[i-4]
	}

	return dp[n]
}

func CountWaysOptimized(n int) int {
	n1, n2, n3, n4 := 1, 1, 1, 2

	for i := 4; i <= n; i++ {
		// todo 1
	}
}
