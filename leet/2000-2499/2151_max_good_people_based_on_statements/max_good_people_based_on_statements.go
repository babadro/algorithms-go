package _2151_max_good_people_based_on_statements

// https://leetcode.com/problems/maximum-good-people-based-on-statements/discuss/1711775/Go-bitmask
// tptl passed #hard
func maximumGood(statements [][]int) int {
	n := len(statements)
	res := 0
Loop:
	for combination := 0; combination < (1 << n); combination++ {
		cnt := 0

		for j := 0; j < n; j++ {
			if isGood(combination, j) {
				cnt++

				for k := 0; k < n; k++ {
					if statements[j][k] == 0 {
						if isGood(combination, k) {
							continue Loop
						}
					} else if statements[j][k] == 1 {
						if !isGood(combination, k) {
							continue Loop
						}
					}
				}
			}
		}

		res = max(res, cnt)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func isGood(num, bit int) bool {
	return num&(1<<bit) != 0
}
