package _350_intersection_of_two_arrays_2

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

func intersectNaive(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int, len(nums1))
	res := make([]int, 0)
	for _, v := range nums1 {
		map1[v]++
	}
	for _, v := range nums2 {
		if map1[v] > 0 {
			res = append(res, v)
			map1[v]--
		}
	}
	return res
}
