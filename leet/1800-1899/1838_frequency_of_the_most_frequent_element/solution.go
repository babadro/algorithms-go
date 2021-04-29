package _1838_frequency_of_the_most_frequent_element

import (
	"github.com/babadro/algorithms-go/utils"
	"sort"
)

// passed. tptl. medium
// https://leetcode.com/problems/frequency-of-the-most-frequent-element/discuss/1178252/Java-or-Sliding-window
func maxFrequency2(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	max, r, l, cum := 0, n-1, n-1, 0
	for l >= 0 && r >= l {
		cum += nums[r] - nums[l]
		if cum <= k {
			l--
			max = utils.Max(max, r-l)
		} else {
			for ; r > 0 && nums[r] == nums[r-1]; r-- {
			}

			cum -= (nums[r] - nums[r-1]) * (r - l)
			r--
			l--
		}
	}

	return max
}
