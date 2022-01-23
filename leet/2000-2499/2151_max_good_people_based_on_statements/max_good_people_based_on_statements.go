package _2151_max_good_people_based_on_statements

// https://leetcode.com/problems/maximum-good-people-based-on-statements/discuss/1711775/Go-bitmask
func maximumGood(statements [][]int) int {
	n := len(statements)
	res := 0
	for i := 0; i < (1 << n); i++ {
		cnt := 0
		good := make([]bool, n)
		for j := 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				cnt++
				good[j] = true
			}
		}

		check := true
		for j := 0; j < n; j++ {
			if good[j] {
				for k := 0; k < n; k++ {
					if statements[j][k] == 0 {
						if good[k] {
							check = false
						}
					}

					if statements[j][k] == 1 {
						if !good[k] {
							check = false
						}
					}
				}
			}
		}

		if check {
			res = max(res, cnt)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
