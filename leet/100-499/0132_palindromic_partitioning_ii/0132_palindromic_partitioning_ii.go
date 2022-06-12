package _132_palindromic_partitioning_ii

// tptl. passed, but slow
// todo 2 implement bottom up
func minCut(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(dp, s, 0, len(s)-1)
}

func recTopDown(dp [][]int, s string, start, end int) int {
	if start >= len(s) || isPalindrome(s, start, end) {
		return 0
	}

	if dp[start][end] == -1 {
		minCuts := end - start
		for i := start; i <= end; i++ {
			if isPalindrome(s, start, i) {
				minCuts = min(minCuts, 1+recTopDown(dp, s, i+1, end))
			}
		}

		dp[start][end] = minCuts
	}

	return dp[start][end]
}

// bruteforce
func minCutBruteForce(s string) int {
	return recBruteForce(s, 0, len(s)-1)
}

func recBruteForce(s string, start, end int) int {
	if start >= len(s) || isPalindrome(s, start, end) {
		return 0
	}

	minCuts := end - start
	for i := start; i <= end; i++ {
		if isPalindrome(s, start, i) {
			minCuts = min(minCuts, 1+recBruteForce(s, i+1, end))
		}
	}

	return minCuts
}

func isPalindrome(s string, start, end int) bool {
	for ; start < end; start, end = start+1, end-1 {
		if s[start] != s[end] {
			return false
		}
	}

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
