package _3_string_interleaving

func findSIIterative(m, n, p string) bool {
	mI, nI, pI := 0, 0, 0
	for (mI < len(m) || nI < len(n)) && pI < len(p) {
		if mI < len(m) && m[mI] == p[pI] {
			mI++
		} else if nI < len(n) && n[nI] == p[pI] {
			nI++
		} else {
			return false
		}

		pI++
	}

	return mI == len(m) && nI == len(n) && pI == len(p)
}
