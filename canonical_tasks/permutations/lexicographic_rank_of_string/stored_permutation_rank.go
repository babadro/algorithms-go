// https://www.geeksforgeeks.org/lexicographic-rank-of-a-string/
// https://www.interviewbit.com/problems/sorted-permutation-rank/
package lexicographic_rank_of_string

const (
	mod     = 1_000_003
	maxChar = 256
)

// todo 1
func findRank(s string) int {
	n := len(s)

}

func factorial(n int) int {
	res := 1

	for i := 1; i <= n; i++ {
		res *= i
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
