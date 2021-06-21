package _1906_minimum_absolute_difference_queries

// tptl. passed. #array, medium (hard for me). best solution
// https://leetcode.com/problems/minimum-absolute-difference-queries/discuss/1284321/Prefix-Sums-vs.-Binary-Search
func minDifference(nums []int, queries [][]int) []int {
	count := make([][101]int, len(nums)+1)

	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= 100; j++ {
			count[i][j] = count[i-1][j]
			if j == nums[i-1] {
				count[i][j]++
			}
		}
	}

	res := make([]int, len(queries))
	for i, q := range queries {
		diff, prev := -1, -1
		for j := 1; j <= 100; j++ {
			if cnt := count[q[1]+1][j] - count[q[0]][j]; cnt > 0 {
				if prev != -1 {
					if diff != -1 {
						diff = min(diff, j-prev)
					} else {
						diff = j - prev
					}
				}

				prev = j
			}
		}

		res[i] = diff
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
