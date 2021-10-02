package _496_next_greater_element_i

// todo 1
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))

Loop:
	for j, num := range nums1 {
		i := 0
		for ; num != nums2[i]; i++ {
		}

		for ; i < len(nums2); i++ {
			if num < nums2[i] {
				res[j] = nums2[i]
				continue Loop
			}
		}

		res[j] = -1
	}

	return res
}
