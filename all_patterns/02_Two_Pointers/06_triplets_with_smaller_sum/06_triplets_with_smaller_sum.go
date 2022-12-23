package _6_triplets_with_smaller_sum

import "sort"

// leetcode (premium) 259 https://leetcode.com/problems/3sum-smaller/
func searchTriplets(arr []int, target int) (res int) {
	sort.Ints(arr)
	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum < target {
				res += right - left
				left++
			} else {
				right--
			}
		}
	}

	return res
}
