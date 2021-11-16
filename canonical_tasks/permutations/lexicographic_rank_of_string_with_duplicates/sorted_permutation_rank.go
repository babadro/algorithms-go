// https://www.interviewbit.com/problems/sorted-permutation-rank-with-repeats/
// https://www.geeksforgeeks.org/lexicographic-rank-string-duplicate-characters/
package lexicographic_rank_of_string_with_duplicates

const (
	mod     = 1_000_003
	maxChar = 256
)

// todo 1
func findRank(s string) int {
	n := len(s)
	tCount := 1

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

		dFac := 1

		for _, ele := range dCount {
			dFac
		}
	}
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
