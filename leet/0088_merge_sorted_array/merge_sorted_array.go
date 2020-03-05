package _088_merge_sorted_array

import "sort"

// best solution
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1[:m+n])
}

// ver 1.0
func merge2(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m)
	copy(temp, nums1)

	for i, j, curr := 0, 0, 0; i < m || j < n; curr++ {
		if i < m && j < n {
			if temp[i] < nums2[j] {
				nums1[curr] = temp[i]
				i++
			} else {
				nums1[curr] = nums2[j]
				j++
			}
		} else if i < m {
			nums1[curr] = temp[i]
			i++
		} else {
			nums1[curr] = nums2[j]
			j++
		}
	}
}
