// leetcode 70
package __staircase

func StairsBruteForce(n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	return StairsBruteForce(n-1) + StairsBruteForce(n-2) + StairsBruteForce(n-3)
}

func StairsTopDown(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}

	return recTopDown(dp, n)
}

func recTopDown(dp []int, n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	if dp[n] == -1 {
		dp[n] = recTopDown(dp, n-1) + recTopDown(dp, n-2) + recTopDown(dp, n-3)
	}

	return dp[n]
}

func StairsBottomUp(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n]
}

func StairsBottomUpOptimized(n int) int {
	n0, n1, n2 := 1, 1, 2
	for i := 3; i <= n; i++ {
		temp := n2 + n1 + n0
		n1, n0 = n2, n1
		n2 = temp
	}

	return n2
}
