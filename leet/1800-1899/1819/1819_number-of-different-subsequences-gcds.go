package _1819

import "github.com/babadro/algorithms-go/utils"

// https://leetcode.com/problems/number-of-different-subsequences-gcds/discuss/1141549/Short-Java-O(NlogN)-Solution-with-explanation
// passed. tptl. hard.
func countDifferentSubsequenceGCDs(nums []int) int {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	set := make([]bool, max+1)
	for _, num := range nums {
		set[num] = true
	}

	ans := 0
	for i := 1; i <= max; i++ {
		d := 0

		for x := i; d != i && x <= max; x += i {
			if set[x] {
				d = utils.GCD(d, x)
			}
		}

		if d == i {
			ans++
		}
	}

	return ans
}
