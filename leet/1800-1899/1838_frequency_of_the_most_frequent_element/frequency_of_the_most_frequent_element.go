package _1838_frequency_of_the_most_frequent_element

import "sort"

const limit = 10_000

// todo 1 doesn't work
// look here: https://leetcode.com/problems/frequency-of-the-most-frequent-element/discuss/1175042/C%2B%2B-Two-Pointers
// or here: https://leetcode.com/problems/frequency-of-the-most-frequent-element/discuss/1178252/Java-or-Sliding-window
func maxFrequency(nums []int, k int) int {
	y := sort.Search(limit, func(i int) bool {
		return sum(i, nums) >= k
	})

	if y == limit {
		return len(nums)
	}

	res := 0
	for _, num := range nums {
		if y >= num {
			res++
		}
	}

	return res
}

func sum(y int, nums []int) int {
	res := 0
	for _, num := range nums {
		if y > num {
			res += y - num
		}
	}

	return res
}
