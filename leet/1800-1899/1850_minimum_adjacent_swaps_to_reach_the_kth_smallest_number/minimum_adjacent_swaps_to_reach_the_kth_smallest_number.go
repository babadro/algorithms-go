package _1850_minimum_adjacent_swaps_to_reach_the_kth_smallest_number

import (
	"fmt"
	"github.com/babadro/algorithms-go/perm"
)

// tptl. passed. medium (hard for me)
// https://leetcode.com/problems/minimum-adjacent-swaps-to-reach-the-kth-smallest-number/discuss/1186921/C%2B%2B-next-permutation
func getMinSwaps(num string, k int) int {
	num1 := []byte(num)
	for i := 0; i < k; i++ {
		_ = perm.NextPermutationBytes(num1)
	}

	var res, j int
	for i := 0; i < len(num); i++ {
		if num[i] != num1[i] {
			for j = i + 1; num[i] != num1[j]; j++ {
			}

			for ; j != i; j-- {
				num1[j], num1[j-1] = num1[j-1], num1[j]

				res++
			}

			fmt.Println(string(num1))
		}
	}

	return res
}
