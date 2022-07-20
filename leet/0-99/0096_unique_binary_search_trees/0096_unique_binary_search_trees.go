package _096_unique_binary_search_trees

// tptl. passed. hard for me
// todo 2 understand
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		j, sum := 1, 0
		for ; j <= i; j++ {
			sum += dp[j-1] * dp[i-j]
		}

		dp[i] = sum
	}

	return dp[n]
}

func rec(start, end int) int {
	if start > end {
		return 1
	}

	res := 0
	for i := start; i <= end; i++ {
		leftCount := rec(i, start-1)
		rightCount := rec(i+1, end)

		res += leftCount * rightCount
	}

	return res
}
