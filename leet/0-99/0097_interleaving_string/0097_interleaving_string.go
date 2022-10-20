package _097_interleaving_string

// tptl. passed
// todo 2 bottom up
func isInterleave(s1 string, s2 string, s3 string) bool {
	return recTopDown(make(map[[3]int]bool), s1, s2, s3, 0, 0, 0)
}

func recTopDown(dp map[[3]int]bool, s1, s2, s3 string, i1, i2, i3 int) bool {
	if i1 == len(s1) && i2 == len(s2) && i3 == len(s3) {
		return true
	}

	if i3 == len(s3) {
		return false
	}

	key := [3]int{i1, i2, i3}
	res, ok := dp[key]
	if !ok {
		res1, res2 := false, false
		if i1 < len(s1) && s1[i1] == s3[i3] {
			res1 = recTopDown(dp, s1, s2, s3, i1+1, i2, i3+1)
		}

		if i2 < len(s2) && s2[i2] == s3[i3] {
			res2 = recTopDown(dp, s1, s2, s3, i1, i2+1, i3+1)
		}

		res = res1 || res2
		dp[key] = res
	}

	return res
}

func isInterleaveBruteForce(s1 string, s2 string, s3 string) bool {
	return recBruteForce(s1, s2, s3, 0, 0, 0)
}

func recBruteForce(s1, s2, s3 string, i1, i2, i3 int) bool {
	if i1 == len(s1) && i2 == len(s2) && i3 == len(s3) {
		return true
	}

	if i3 == len(s3) {
		return false
	}

	res1, res2 := false, false
	if i1 < len(s1) && s1[i1] == s3[i3] {
		res1 = recBruteForce(s1, s2, s3, i1+1, i2, i3+1)
	}

	if i2 < len(s2) && s2[i2] == s3[i3] {
		res2 = recBruteForce(s1, s2, s3, i1, i2+1, i3+1)
	}

	return res1 || res2
}
