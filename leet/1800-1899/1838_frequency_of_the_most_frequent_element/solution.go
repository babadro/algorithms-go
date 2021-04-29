package _1838_frequency_of_the_most_frequent_element

import (
	"github.com/babadro/algorithms-go/utils"
	"sort"
)

// passed. tptl. medium (hard for me)
// https://leetcode.com/problems/frequency-of-the-most-frequent-element/discuss/1175042/C%2B%2B-Two-Pointers
// todo 1 need to understand and outline
func maxFrequency2(nums []int, k int) int {
	res := 1
	sort.Ints(nums)
	i := len(nums) - 1
	for j := i; i > 0; i-- {
		for ; j >= 0 && k >= nums[i]-nums[j]; j-- {
			k -= nums[i] - nums[j]
		}

		res = utils.Max(res, i-j)
		k += (nums[i] - nums[i-1]) * (i - j - 1)
	}

	return res
}
