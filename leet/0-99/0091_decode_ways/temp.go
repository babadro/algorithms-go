package _091_decode_ways

func numDecodings2(s string) int {
	dp := make([]int, len(s))
	for i := range dp {
		dp[i] = -1
	}

	return rec(dp, s, 0, 0)
}

func rec(dp []int, s string, i int, prev byte) int {
	num := prev*10 + (s[i] - '0')

	if num == 0 || num > 26 {
		return 0
	}

	if i == len(s)-1 {
		return 1
	}

	if dp[i] == -1 {
		dp[i] = rec(dp, s, i+1, 0) + rec(dp, s, i+1, s[i]-'0')
	}

	return dp[i]
}

// bruteforce. tle
func numDecodings3(s string) int {
	var res int
	rec3(s, 0, 0, &res)

	return res
}

func rec3(s string, i int, prev byte, res *int) {
	num := prev*10 + (s[i] - '0')

	if num == 0 || num > 26 {
		return
	}

	if i == len(s)-1 {
		*res++
		return
	}

	rec3(s, i+1, 0, res)
	if prev == 0 {
		rec3(s, i+1, s[i]-'0', res)
	}
}
