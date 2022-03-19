package _2122_recover_the_original_array

import "sort"

// todo 1 doesn't work. Try C++ solution in comments
// https://leetcode.com/problems/recover-the-original-array/discuss/1647452/Python-Short-solution-explained
func recoverArray(nums []int) []int {
	n := len(nums) / 2
	ans := make([]int, n)

	sort.Ints(nums)
	m := make(map[int]int, n)

	for _, num := range nums {
		m[num] = m[num] + 1
	}

	for i := 1; i <= n; i++ {
		k := nums[i] - nums[0]
		if k != 0 && k%2 == 0 {
			if check(nums, ans, copyMap(m), k) {
				return ans
			}
		}
	}

	return ans
}

func check(nums, ans []int, m map[int]int, k int) bool {
	idx := 0
	for _, num := range nums {
		if m[num] == 0 {
			continue
		}

		if m[num+k] == 0 {
			return false
		}

		m[num] = m[num] - 1
		m[num] = m[num+k] - 1

		ans[idx] = num + k/2
		idx++
	}

	return true
}

func copyMap(m map[int]int) map[int]int {
	res := make(map[int]int, len(m))
	for k, v := range m {
		res[k] = v
	}

	return res
}
