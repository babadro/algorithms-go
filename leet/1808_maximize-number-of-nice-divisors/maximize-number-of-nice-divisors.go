package _1808_maximize_number_of_nice_divisors

// todo 1 implement pure math solution without st
// https://leetcode.com/problems/maximize-number-of-nice-divisors/discuss/1130586/C%2B%2BJava-modpow
func maxNiceDivisors(n int) int {
	st := []int{0, 1, 2, 3, 4, 6}
	mod := 1_000_000_007

	if n < 6 {
		return st[n]
	}

	return (ModPow(3, n/3-1, mod) * st[3+n%3]) % mod
}

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

		base = (base * base) % modulus
	}

	return result
}

func maxNiceDivisorsBruteForce(n int) int {
	if n == 1 {
		return 1
	}

	max := 0
	for count := 1; count <= n; count++ {
		num := n / count

		remainder := n % count

		arr := make([]int, count)
		for i := range arr {
			arr[i] = num
		}

		for i := 0; remainder > 0; i++ {
			i %= len(arr)
			arr[i]++
			remainder--
		}

		product := 1
		for _, number := range arr {
			product *= number
		}

		if product > max {
			max = product
		}
	}

	return max
}
