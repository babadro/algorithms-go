package _1850_minimum_adjacent_swaps_to_reach_the_kth_smallest_number

import (
	"github.com/babadro/algorithms-go/perm"
)

// todo 1:
// https://leetcode.com/problems/minimum-adjacent-swaps-to-reach-the-kth-smallest-number/discuss/1186921/C%2B%2B-next-permutation
func getMinSwaps(num string, k int) int {
	num1 := []byte(num)
	for i := 0; i < k; i++ {
		perm.NextPermutationBytes(num1)
	}

	res, n := 0, len(num)
	for i := 0; i < n; i++ {
		if num[i] != num1[i] {
			for j := i + 1; j < n; j++ {
				if num[i] == num1[j] {
					res += j - i

					str := string(num1[:i+1]) + string(num1[i:j]) + string(num1[j+1:])
					num1 = []byte(str)
					//num1 = append(num1[:i+1], num1[i:j]...)
					//num1 = append(num1, num1[j+1:]...)

					//fmt.Println(string(num1))
					break
				}
			}
		}
	}

	return res
}
