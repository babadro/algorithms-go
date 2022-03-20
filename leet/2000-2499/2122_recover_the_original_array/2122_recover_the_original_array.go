package _2122_recover_the_original_array

import "sort"

// passed. ntu idea
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
		m[num+k] = m[num+k] - 1

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

// 4 times faster and pretty. #ntu
func recoverArray2(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums)/2; i++ {
		if k := nums[i+1] - nums[0]; k > 0 && k%2 == 0 {
			var res, deque []int
			for _, v := range nums {
				if len(deque) > 0 && deque[0] == v {
					deque = deque[1:]
					res = append(res, v-k/2)
				} else {
					deque = append(deque, v+k)
				}
			}

			if len(deque) == 0 {
				return res
			}
		}
	}

	return nil
}

// 20 times faster. best solution. #ntu
func recoverArray3(nums []int) []int {
	n := len(nums)
	sort.Ints(nums)

	for i := 1; i < n; i++ {
		res := check3(i, n, nums)

		if len(res) == n/2 {
			return res
		}
	}

	return nil
}

func check3(i, n int, nums []int) []int {
	k := nums[i] - nums[0]
	if k%2 != 0 {
		return nil
	}

	used := make([]bool, n)
	var res []int
	j := 0
	for j < n && i < n {
		if used[j] {
			j++
			continue
		}

		t := nums[j] + k
		if t == nums[i] {
			res = append(res, nums[j]+k/2)
			used[i], used[j] = true, true
			i++
			j++
		} else if t < nums[i] {
			return nil
		} else {
			i++
		}
	}

	return res
}
