package _1855_maximum_distance_between_a_pair_of_values

import "sort"

// tptl. passed. medium. best solution
func maxDistance2(nums1 []int, nums2 []int) int {
	max := 0
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if i > j {
			j++
			continue
		}

		if nums1[i] > nums2[j] {
			i++
			continue
		}

		if dist := j - i; dist > max {
			max = dist
		}

		j++
	}

	return max
}

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
