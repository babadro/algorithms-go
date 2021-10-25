package __fibonacci_numbers

func CalculateFibonacci(n int) int {
	if n < 2 {
		return n
	}

	return CalculateFibonacci(n-2) + CalculateFibonacci(n-1)
}

func CalculateFibonacciTopDown(n int) int {
	dp := make([]int, n+1)

	return topDownRec(dp, n)
}

func topDownRec(dp []int, n int) int {
	if n < 2 {
		return n
	}

	if dp[n] == 0 {
		dp[n] = topDownRec(dp, n-1) + topDownRec(dp, n-2)
	}

	return dp[n]
}

func CalculateFibonacciBottomUp(n int) int {
	if n < 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func CalculateFibonacciOptimized(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 0, 1
	for i := 2; i <= n; i++ {
		n2, n1 = n2+n1, n2
	}

	return n2
}
