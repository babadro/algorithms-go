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
