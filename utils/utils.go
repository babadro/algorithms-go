package utils

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Max3(a, b, c int) int {
	res := a
	if b > res {
		res = b
	}

	if c > res {
		res = c
	}

	return res
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Min3(a, b, c int) int {
	res := a
	if b < res {
		res = b
	}

	if c < res {
		res = c
	}

	return res
}

func Abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}
