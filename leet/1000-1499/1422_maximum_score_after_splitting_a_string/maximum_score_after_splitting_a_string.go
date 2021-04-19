package _1422_maximum_score_after_splitting_a_string

// passed. best solution. tptl
// https://leetcode.com/problems/maximum-score-after-splitting-a-string/discuss/601297/Go-very-simple-Solution
func maxScore2(s string) int {
	count0, count1 := 0, 0
	for _, char := range s {
		if char == '0' {
			count0++
		} else {
			count1++
		}
	}

	max := 0
	current := count1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			current++
		} else {
			current--
		}

		if current > max {
			max = current
		}
	}

	return max
}

// passed, but two slices demands memory
func maxScore(s string) int {
	n := len(s)
	zeroesSum, onesSum := make([]int, n), make([]int, n)

	counter := 0
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			counter++
			zeroesSum[i] = counter
		}
	}

	counter = 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' {
			counter++
			onesSum[i] = counter
		}
	}

	max := 0
	for i := 0; i < n-1; i++ {
		score := zeroesSum[i] + onesSum[i+1]
		if score > max {
			max = score
		}
	}

	return max
}
