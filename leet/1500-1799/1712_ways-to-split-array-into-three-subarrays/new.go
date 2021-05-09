package _1712_ways_to_split_array_into_three_subarrays

import (
	"sort"
)

func waysToSplit3(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}

	var left, mid, right int
	for k := n - 1; k >= 2; k-- {
		end := sort.Search(k, func(i int) bool {
			left, mid, right = nums[i], nums[k-1]-nums[i], nums[n-1]-nums[k-1]

			return left >= mid && mid <= right
		})

		if left > mid || mid < right {
			end--
		}

		//if count := end

	}

	return 0
}

/*
end := sort.Search(k, func(i int) bool {
			left, mid, right := nums[i], nums[k-1]-nums[i], nums[n-1]-nums[k-1]

			return left >= mid && mid <= right
		})

		left, mid, right := nums[end], nums[k-1]-nums[end], nums[n-1]-nums[k-1]
*/

/*
	j := sort.Search(k-1, func(i int) bool {
			return nums[k-1]-nums[i] <= nums[i]
		})

		if nums[k-1]-nums[j] < nums[j] {
			j--
		}

		m := sort.Search(k-1, func(i int) bool {
			return nums[k-1]-nums[i] <= nums[n-1]-nums[k-1]
		})

		if nums[k-1]-nums[m] < nums[n-1]-nums[k-1] {
			m--
		}

		idx := utils.Min(j, m)
		if idx < 0 {
			break
		}
*/
