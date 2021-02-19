package _1712_ways_to_split_array_into_three_subarrays

import "sort"

func waysToSplit(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}

	counter := 0
	for i := 0; i < n-2; i++ {
		//if nums[i] >= nums[n-1]-nums[i] {
		//	break
		//}

		start := sort.Search(n, func(j int) bool {
			return nums[j]-nums[i] >= nums[i]
		})

		finish := sort.Search(n, func(j int) bool {
			return nums[j]-nums[i] > nums[n-1]-nums[j]
		})

		//var idx int
		//for j := i; j < n; j++ {
		//	if nums[j]-nums[i] > nums[n-1]-nums[j] {
		//		idx = j
		//		break
		//	}
		//}

		count := finish - start
		if count <= 0 {
			break
		}

		counter += count
	}

	return counter % 1_000_000_007
}
