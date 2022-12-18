package _070_climbing_stairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	prev := 2
	prevPrev := 1
	for i := 2; i < n; i++ {
		sum := prev + prevPrev
		prevPrev = prev
		prev = sum
	}
	return prev
}

func climbStairsBruteForce(n int) int {
	if n == 2 {
		return 2
	}

	if n == 1 {
		return 1
	}

	return climbStairsBruteForce(n-1) + climbStairsBruteForce(n-2)
}

func climbStairsTopDown(n int) int {
	return recTopDown(make([]int, n+1), n)
}

func recTopDown(dp []int, n int) int {
	if n == 2 {
		return 2
	}

	if n == 1 {
		return 1
	}

	if dp[n] == 0 {
		dp[n] = recTopDown(dp, n-1) + recTopDown(dp, n-2)
	}

	return dp[n]
}
