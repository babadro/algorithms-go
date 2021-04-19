package _1759_count_number_of_homogenous_substrings

// passed. tptl
func countHomogenous(s string) int {
	curr, res := 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curr++
		} else {
			res += sumArithmeticProgression(curr)
			curr = 1
		}
	}

	return (res + sumArithmeticProgression(curr)) % 1000000007
}

func sumArithmeticProgression(n int) int {
	return (1 + n) * n / 2
}
