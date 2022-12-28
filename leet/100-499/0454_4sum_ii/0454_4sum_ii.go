package _454_4sum_ii

import "sort"

// tptl. passed. not two pointer
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			m[num1+num2]++
		}
	}

	res := 0
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			res += m[-(num3 + num4)]
		}
	}

	return res
}

// tptl. passed. very slow, but implements two pointers approach
// todo2 implement approach with map[sum]count where sum = nums1[i] + nums2[j] for all i and j
func fourSumCount2(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sort.Ints(nums3)
	sort.Ints(nums4)

	n := len(nums1)
	var res int
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			left, right := 0, n-1
			for left < n && right >= 0 {
				if sum := num1 + num2 + nums3[left] + nums4[right]; sum < 0 {
					left++
				} else if sum > 0 {
					right--
				} else {
					tmpLeft, tmpRight := left, right
					for tmpLeft < n && nums3[tmpLeft] == nums3[left] {
						tmpLeft++
					}

					for tmpRight >= 0 && nums4[tmpRight] == nums4[right] {
						tmpRight--
					}

					res += (tmpLeft - left) * (right - tmpRight)

					left, right = tmpLeft, tmpRight
				}
			}
		}
	}

	return res
}

// bruteforce, TLE
func fourSumCountTLE(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sort.Ints(nums3)
	sort.Ints(nums4)

	var res int
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
		Loop:
			for _, num3 := range nums3 {
				for j := len(nums4) - 1; j >= 0; j-- {
					sum := num1 + num2 + num3 + nums4[j]
					if sum < 0 {
						continue Loop
					} else if sum > 0 {
						continue
					} else {
						res++
					}
				}
			}
		}
	}

	return res
}
