package _35

import "github.com/babadro/algorithms-go/utils"

// 1819 todo 1 https://leetcode.com/problems/number-of-different-subsequences-gcds/discuss/1141549/Short-Java-O(NlogN)-Solution-with-explanation
func countDifferentSubsequenceGCDs(nums []int) int {
	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	mem := make([]bool, maxNum+1)
	for _, num := range nums {
		mem[num] = true
	}

	ans := 0
	for i := 1; i <= maxNum; i++ {
		g := -1
		for j := 1; j < maxNum; j += i {
			if mem[j] {
				if g == -1 {
					g = j
				} else {
					g = utils.GCD(g, j)
				}
			}

			if g == i {
				ans++
			}
		}
	}

	return ans
}
