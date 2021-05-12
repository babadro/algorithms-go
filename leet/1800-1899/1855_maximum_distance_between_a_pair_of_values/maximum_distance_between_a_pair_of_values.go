package _855_maximum_distance_between_a_pair_of_values

import "sort"

// passed
func maxDistance(nums1 []int, nums2 []int) int {
	max := 0
	for j := len(nums2) - 1; j >= 0; j-- {
		idx := sort.Search(len(nums1), func(i int) bool {
			return nums1[i] <= nums2[j]
		})

		if idx == len(nums1) || idx > j {
			continue
		}

		if dist := j - idx; dist > max {
			max = dist
		}
	}

	return max
}
