package _1712_ways_to_split_array_into_three_subarrays

import "sort"

func waysToSplit4(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}

	counter := 0
	for k := n - 2; k >= 1; k-- {
		start := sort.Search(k, func(i int) bool {
			r, m := nums[n-1]-nums[k], nums[k]-nums[i]

			return r >= m
		})

		finish := sort.Search(k, func(i int) bool {
			m, l := nums[k]-nums[i], nums[i]

			return m >= l
		})

		if start >= 1 && finish >= start {
			counter += finish - start + 1
		}
	}

	return counter % 1_000_000_007
}
