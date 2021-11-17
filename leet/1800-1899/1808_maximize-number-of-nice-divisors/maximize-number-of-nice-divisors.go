// https://leetcode.com/problems/minimum-number-of-operations-to-make-string-sorted/discuss/1163050/Python-O(26n)-math-solution-explained
package _1808_maximize_number_of_nice_divisors

// tptl. passed. best solution. hard
func maxNiceDivisors2(n int) int {
	if n <= 4 {
		return n
	}

	remainder, quotient, mod := n%3, n/3, 1_000_000_007
	if remainder == 0 {
		return ModPow(3, quotient, mod)
	}

	if remainder == 1 {
		return (4 * ModPow(3, quotient-1, mod)) % mod
	}

	return (2 * ModPow(3, quotient, mod)) % mod
}

// https://leetcode.com/problems/maximize-number-of-nice-divisors/discuss/1130586/C%2B%2BJava-modpow
func maxNiceDivisors(n int) int {
	st := []int{0, 1, 2, 3, 4, 6}
	mod := 1_000_000_007

	if n < 6 {
		return st[n]
	}

	idx := 3 + n%3
	exp := n/3 - 1
	_, _ = idx, exp
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

func maxNiceDivisorsBruteForce(n int) (int, []int) {
	if n == 1 {
		return 1, []int{1}
	}

	max, bestArr := 0, []int(nil)
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
			bestArr = arr
		}
	}

	return max, bestArr
}
