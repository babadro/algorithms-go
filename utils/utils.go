package utils

import "math"

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

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Make2DArr(m, n int, defaultVal int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		if defaultVal != 0 {
			for j := range res[i] {
				res[i][j] = defaultVal
			}
		}
	}

	return res
}

func Make3DArr(d1, d2, d3 int, defaultVal int) [][][]int {
	res := make([][][]int, d1)
	for i := range res {
		res[i] = make([][]int, d2)
		for j := range res[i] {
			res[i][j] = make([]int, d3)
			if defaultVal != 0 {
				for k := range res[i][j] {
					res[i][j][k] = defaultVal
				}
			}
		}

	}

	return res
}

// tptl
func ModPow(base, exp, modulus int) int {
	result := 1
	base %= modulus
	if base == 0 {
		return 0
	}

	for ; exp > 0; exp >>= 1 {
		if (exp & 1) == 1 {
			result = (result * base) % modulus
		}

		base = (base * base) % modulus // because a^(m*n) = (a^m)^n
	}

	return result
}

// tptl
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

// tptl
func Log(b, x int) (pow int) {
	for {
		if x <= b-1 {
			return pow
		}

		pow++
		x /= b
	}
}
