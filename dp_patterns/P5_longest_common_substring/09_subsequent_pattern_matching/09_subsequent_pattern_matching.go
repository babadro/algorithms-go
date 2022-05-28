package _9_subsequent_pattern_matching

func findSPMCountBruteForce(st, pat string) int {
	return recBruteForce(st, pat, 0, 0)
}

func recBruteForce(st, pat string, i1, i2 int) int {
	if i2 == len(pat) {
		return 1
	}

	if i1 == len(st) {
		return 0
	}

	res := 0
	if st[i1] == pat[i2] {
		res += recBruteForce(st, pat, i1+1, i2+1)
	}

	res += recBruteForce(st, pat, i1+1, i2)

	return res
}

func findSPMCountTopDown(st, pat string) int {
	dp := make([][]int, len(st))
	for i := range dp {
		dp[i] = make([]int, len(pat))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(dp, st, pat, 0, 0)
}

func recTopDown(dp [][]int, st, pat string, i1, i2 int) int {
	if i2 == len(pat) {
		return 1
	}

	if i1 == len(st) {
		return 0
	}

	if dp[i1][i2] == -1 {
		res := 0
		if st[i1] == pat[i2] {
			res += recTopDown(dp, st, pat, i1+1, i2+1)
		}

		res += recTopDown(dp, st, pat, i1+1, i2)

		dp[i1][i2] = res
	}

	return dp[i1][i2]
}

func findSPMCountBottomUp(st, pat string) int {
	dp := make([][]int, len(st)+1)
	for i := range dp {
		dp[i] = make([]int, len(pat)+1)
	}

	// for the empty pattern we have one matching
	for i := 0; i < len(st); i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= len(st); i++ {
		for j := 1; j <= len(pat); j++ {
			if st[i-1] == pat[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}

			dp[i][j] += dp[i-1][j]
		}
	}

	return dp[len(st)][len(pat)]
}
