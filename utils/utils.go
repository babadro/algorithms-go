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

func Min4(a, b, c, d int) int {
	res := a
	if b < res {
		res = b
	}

	if c < res {
		res = c
	}

	if d < res {
		res = d
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

func Ptr[T any](v T) *T {
	return &v
}

const mod = 1_000_000_007

func PowerMod(x, y int) int {
	res := 1
	base := x % mod

	for y > 0 {
		if y%2 == 1 {
			res = res * base % mod
		}

		base = base * base % mod

		y /= 2
	}

	return res
}

func FactorialMod(n int) int {
	res := 1

	for i := 2; i <= n; i++ {
		res = res * i % mod
	}

	return res
}

// n!/((n-k)!*k!)
func CombinationNOverKMod(n, k int) int {
	if k > n {
		return 0
	}

	num := FactorialMod(n)

	den := FactorialMod(k) * FactorialMod(n-k) % mod

	return num * PowerMod(den, mod-2) % mod
}
