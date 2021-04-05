package _35

import "github.com/babadro/algorithms-go/utils"

// 1819 todo 1
func countDifferentSubsequenceGCDs(nums []int) int {
	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	flags := make([]bool, maxNum+1)
	for _, num := range nums {
		flags[num] = true
	}

	res := 0
	for i := 1; i <= maxNum; i++ {
		var niGcd int
		if flags[i] {
			niGcd = i
		} else {
			niGcd = 0
		}

		for n := 1; n*i <= maxNum && niGcd != i; n++ {
			if flags[n*i] {
				if niGcd != 0 {
					niGcd = utils.GCD(niGcd, n*i)
				} else {
					niGcd = n * i
				}
			}

			if niGcd == i {
				res += niGcd
			}
		}
	}

	return res
}
