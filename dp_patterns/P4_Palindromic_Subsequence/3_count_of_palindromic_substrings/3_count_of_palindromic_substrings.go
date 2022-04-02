package __count_of_palindromic_substrings

func findCPSBottomUp(st string) int {
	dp := make([][]bool, len(st))
	for i := range dp {
		dp[i] = make([]bool, len(st))
	}

	count := 0
	for i := 0; i < len(st); i++ {
		dp[i][i] = true
		count++
	}

	for startIDx := len(st) - 1; startIDx >= 0; startIDx-- {
		for endIDx := startIDx + 1; endIDx < len(st); endIDx++ {
			if st[startIDx] == st[endIDx] {
				if endIDx-startIDx == 1 || dp[startIDx+1][endIDx-1] {
					dp[startIDx][endIDx] = true
					count++
				}
			}
		}
	}

	return count
}
