package _091_decode_ways

func numDecodings2(s string) int {
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, len(s))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return rec(dp, s, 0, false)
}

func rec(dp [][]int, s string, i int, prev bool) int {
	prevNum := byte(0)
	if prev {
		prevNum = (s[i-1] - '0') * 10
	}

	num := prevNum + (s[i] - '0')

	if num == 0 || num > 26 {
		return 0
	}

	if i == len(s)-1 {
		return 1
	}

	idx := 0
	if prev {
		idx = 1
	}

	if dp[idx][i] == -1 {
		dp[idx][i] = rec(dp, s, i+1, false) + rec(dp, s, i+1, true)
	}

	return dp[idx][i]
}

// bruteforce. tle
func numDecodings3(s string) int {
	return rec3(s, 0, 0)
}

func rec3(s string, i int, prev byte) int {
	num := prev*10 + (s[i] - '0')

	if num == 0 || num > 26 {
		return 0
	}

	if i == len(s)-1 {
		return 1
	}

	res := rec3(s, i+1, 0)
	if prev == 0 {
		res += rec3(s, i+1, s[i]-'0')
	}

	return res
}