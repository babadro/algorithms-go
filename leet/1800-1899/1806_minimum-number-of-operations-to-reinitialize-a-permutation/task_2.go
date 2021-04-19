package _1806_minimum_number_of_operations_to_reinitialize_a_permutation

import "github.com/babadro/algorithms-go/slices"

// tptl. passed. bruteforce. hard. todo 2 read about more effective solution https://leetcode.com/problems/minimum-number-of-operations-to-reinitialize-a-permutation/discuss/1132271/C%2B%2BJava-Easy-to-understand-solution-w-explanation-faster-than-100.00
func reinitializePermutation(n int) int {
	arr, tmp1, tmp2 := make([]int, n), make([]int, n), make([]int, n)
	for i := range arr {
		arr[i] = i
		tmp1[i] = i
	}

	count := 0

	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				tmp2[i] = tmp1[i/2]
			} else {
				tmp2[i] = tmp1[n/2+(i-1)/2]
			}
		}

		count++
		if slices.IntSlicesAreEqual(tmp2, arr) {
			return count
		}

		tmp1, tmp2 = tmp2, tmp1
	}

	return -1
}
