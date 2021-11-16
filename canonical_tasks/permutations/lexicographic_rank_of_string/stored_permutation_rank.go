// https://www.geeksforgeeks.org/lexicographic-rank-of-a-string/
// https://www.interviewbit.com/problems/sorted-permutation-rank/
package lexicographic_rank_of_string

const (
	mod     = 1_000_003
	maxChar = 256
)

// passed. tptl
func findRank(s string) int {
	n := len(s)

	rank := 1

	count := populateCount(s)
	factorialMemo := modFactorial(n, mod)

	for i := 0; i < n; i++ {
		rank += count[s[i]-1] * factorialMemo[n-1-i]

		rank %= mod

		count = updateCount(count, int(s[i]))
	}

	return rank
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

func populateCount(str string) [maxChar]int {
	count := [maxChar]int{}
	for i := 0; i < len(str); i++ {
		count[str[i]]++
	}

	for i := 1; i < maxChar; i++ {
		count[i] += count[i-1]
	}

	return count
}

func updateCount(count [maxChar]int, char int) [maxChar]int {
	for i := char; i < maxChar; i++ {
		count[i]--
	}

	return count
}
