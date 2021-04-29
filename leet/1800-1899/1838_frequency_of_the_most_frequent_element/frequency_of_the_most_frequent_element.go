package _1838_frequency_of_the_most_frequent_element

import "sort"

const limit = 10_000

// doesn't work
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
