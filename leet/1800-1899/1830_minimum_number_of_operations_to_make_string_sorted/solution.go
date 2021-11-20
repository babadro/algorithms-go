package _1830_minimum_number_of_operations_to_make_string_sorted

const (
	mod = 1_000_000_007
)

// tptl passed, but slow (244 ms). Todo 1 - optimize it
func makeStringSorted(s string) int {
	n := len(s)
	var ans int
	factMemo := modFactorial(n, mod)

	for i := 0; i < n; i++ {
		lessThan := 0
		for j := i + 1; j < n; j++ {
			if s[i] > s[j] {
				lessThan++
			}
		}

		dCount := [26]int{}
		for j := i; j < n; j++ {
			dCount[s[j]-'a']++
		}

		denominator := 1
		for _, ele := range dCount {
			denominator = (denominator * factMemo[ele]) % mod
		}

		inverse := modPow(denominator, mod-2, mod)

		permutation := (factMemo[n-i-1] * inverse) % mod
		totalPermutation := (lessThan * permutation) % mod

		ans += totalPermutation
		ans %= mod
	}

	return ans
}

func modFactorial(n, mod int) []int {
	if n == 0 {
		return nil
	}

	res := make([]int, n+1)
	res[0] = 1

	for i := 1; i <= n; i++ {
		res[i] = i * res[i-1]
		res[i] %= mod
	}

	return res
}

func modPow(base, exp, modulus int) int {
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
