package _3_string_interleaving

func findSIBruteForce(m, n, p string) bool {
	return recBruteForce(m, n, p, 0, 0, 0)
}

func recBruteForce(m, n, p string, mIDx, nIDx, pIDx int) bool {
	if mIDx == len(m) && nIDx == len(n) && pIDx == len(p) {
		return true
	}

	if pIDx == len(p) {
		return false
	}

	b1, b2 := false, false
	if mIDx < len(m) && m[mIDx] == p[pIDx] {
		b1 = recBruteForce(m, n, p, mIDx+1, nIDx, pIDx+1)
	}

	if nIDx < len(n) && n[nIDx] == p[pIDx] {
		b2 = recBruteForce(m, n, p, mIDx, nIDx+1, pIDx+1)
	}

	return b1 || b2
}

func findSITopDown(m, n, p string) bool {
	return recTopDown(make(map[[3]int]bool), m, n, p, 0, 0, 0)
}

func recTopDown(dp map[[3]int]bool, m, n, p string, mIDx, nIDx, pIDx int) bool {
	if mIDx == len(m) && nIDx == len(n) && pIDx == len(p) {
		return true
	}

	if pIDx == len(p) {
		return false
	}

	key := [3]int{mIDx, nIDx, pIDx}
	if _, ok := dp[key]; !ok {
		b1, b2 := false, false
		if mIDx < len(m) && m[mIDx] == p[pIDx] {
			b1 = recTopDown(dp, m, n, p, mIDx+1, nIDx, pIDx+1)
		}

		if nIDx < len(n) && n[nIDx] == p[pIDx] {
			b2 = recTopDown(dp, m, n, p, mIDx, nIDx+1, pIDx+1)
		}

		dp[key] = b1 || b2
	}

	return dp[key]
}

func findSIBottomUp(m, n, p string) bool {
	if len(m)+len(n) != len(p) {
		return false
	}

	dp := make([][]bool, len(m)+1)
	for i := range dp {
		dp[i] = make([]bool, len(n)+1)
	}

	for mIDx := 0; mIDx <= len(m); mIDx++ {
		for nIDx := 0; nIDx <= len(n); nIDx++ {
			// if 'm' and 'n' are empty, then 'p' must have been empty too.
			if mIDx == 0 && nIDx == 0 {
				dp[mIDx][nIDx] = true
				// if 'm' is empty, we need to check the interleaving with 'n' only
			} else if mIDx == 0 && n[nIDx-1] == p[mIDx+nIDx-1] {
				dp[mIDx][nIDx] = dp[mIDx][nIDx-1]
				// if 'n' is empty, we need to check the interleaving with 'm' only
			} else if nIDx == 0 && m[mIDx-1] == p[mIDx+nIDx-1] {
				dp[mIDx][nIDx] = dp[mIDx-1][nIDx]
			} else {
				b1, b2 := false, false
				// if the letter of 'm' and 'p' match, we take whatever is matched till mIndex-1
				if mIDx > 0 && m[mIDx-1] == p[mIDx+nIDx-1] {
					b1 = dp[mIDx-1][nIDx]
				}
				// if the letter of 'n' and 'p' match, we take whatever is matched till nIndex-1 too
				// note the '|=', this is required when we have common letters
				if nIDx > 0 && n[nIDx-1] == p[mIDx+nIDx-1] {
					b2 = dp[mIDx][nIDx-1]
				}

				dp[mIDx][nIDx] = b1 || b2
			}
		}
	}

	return dp[len(m)][len(n)]
}
