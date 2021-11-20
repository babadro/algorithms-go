// https://www.interviewbit.com/problems/sorted-permutation-rank-with-repeats/
// https://www.geeksforgeeks.org/lexicographic-rank-string-duplicate-characters/
// https://www.algostreak.com/post/sorted-permutation-rank-with-repeats-interviewbit-solution
package lexicographic_rank_of_string_with_duplicates

const (
	mod = 1_000_003
)

// tptl passed
func findRank(s string) int {
	n := len(s)
	ans := 1
	factMemo := modFactorial(n, mod)

	for i := 0; i < n; i++ {
		lessThan := 0
		for j := i + 1; j < n; j++ {
			if s[i] > s[j] {
				lessThan++
			}
		}

		dCount := [256]int{}
		for j := i; j < n; j++ {
			dCount[s[j]]++
		}

		denominator := 1
		for _, ele := range dCount {
			denominator = (denominator * factMemo[ele]) % mod
		}

		inverse := ModPow(denominator, mod-2, mod)

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
