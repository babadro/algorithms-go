package _091_decode_ways

// dp top down. tptl. passed
func numDecodings3(s string) int {
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, len(s))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return rec3(dp, s, 0, false)
}

func rec3(dp [][]int, s string, i int, takePrev bool) int {
	num := byte(0)
	if takePrev {
		num = (s[i-1] - '0') * 10
	}

	num += s[i] - '0'

	if num == 0 || num > 26 {
		return 0
	}

	if i == len(s)-1 {
		return 1
	}

	dpIDx := 0
	if takePrev {
		dpIDx = 1
	}

	if dp[dpIDx][i] == -1 {
		res := rec3(dp, s, i+1, false)
		if !takePrev {
			res += rec3(dp, s, i+1, true)
		}

		dp[dpIDx][i] = res
	}

	return dp[dpIDx][i]
}

// bruteforce. tle
func numDecodings2(s string) int {
	return rec2(s, 0, 0)
}

func rec2(s string, i int, prev byte) int {
	num := prev*10 + (s[i] - '0')

	if num == 0 || num > 26 {
		return 0
	}

	if i == len(s)-1 {
		return 1
	}

	res := rec2(s, i+1, 0)
	if prev == 0 {
		res += rec2(s, i+1, s[i]-'0')
	}

	return res
}

// 100%, 93%. dp bottom up
func numDecodings(s string) int {
	n := len(s)
	if s[0] == '0' {
		return 0
	}

	arr := make([]int, n)
	if last := n - 1; s[last] != '0' {
		arr[last] = 1
	}

	for i := n - 2; i >= 0; i-- {
		if nextNum := arr[i+1]; nextNum == 0 {
			arr[i] = 1
		} else {
			arr[i] = nextNum
		}

		currChar, nextChar := s[i], s[i+1]
		if nextChar == '0' {
			if currChar == '0' || currChar > '2' {
				return 0
			}

			continue
		}

		idx := i + 2
		if (currChar == '1' && nextChar > '0') || (currChar == '2' && nextChar > '0' && nextChar <= '6') {
			if idx >= n {
				arr[i] = 2
			} else {
				arr[i] += arr[idx]
			}
		}
	}

	return arr[0]
}
