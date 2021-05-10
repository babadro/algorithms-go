package _1856_maximum_subarray_min_product

import "github.com/babadro/algorithms-go/utils"

// tptl. medium (hard for me). passed. best solution
// https://leetcode.com/problems/maximum-subarray-min-product/discuss/1199242/Java-Simple-Stack-Solution
// todo 2 need to understand (or find easier solution)
func maxSumMinProduct3(nums []int) int {
	st, cnt := make([]int, 0), make([]int, 0)
	max, c := 0, 0

	for i := 0; i < len(nums); i++ {
		c = nums[i]
		for len(st) != 0 && st[len(st)-1] >= nums[i] {
			c += cnt[len(cnt)-1]
			cnt = cnt[:len(cnt)-1]

			max = utils.Max(max, (c-nums[i])*st[len(st)-1])
			st = st[:len(st)-1]
		}

		st = append(st, nums[i])
		cnt = append(cnt, c)
		max = utils.Max(max, c*nums[i])
	}

	c = 0
	for len(st) != 0 {
		c += cnt[len(cnt)-1]
		cnt = cnt[:len(cnt)-1]

		max = utils.Max(max, c*st[len(st)-1])
		st = st[:len(st)-1]
	}

	return max % 1_000_000_007
}

// passed. effective but hard to understand
// https://leetcode.com/problems/maximum-subarray-min-product/discuss/1198896/O(n)-Monostack-with-picture
func maxSumMinProduct2(nums []int) int {
	nums = append(nums, 0)
	n := len(nums)

	dp, st := make([]int, n+1), make([]int, 0)
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i] + nums[i]
	}

	res := 0
	for i := 0; i < n; i++ {
		for len(st) != 0 && nums[st[len(st)-1]] > nums[i] {
			j := st[len(st)-1]
			st = st[:len(st)-1]

			idx := 0
			if len(st) > 0 {
				idx = st[len(st)-1] + 1
			}

			if curr := nums[j] * (dp[i] - dp[idx]); curr > res {
				res = curr
			}
		}

		st = append(st, i)
	}

	return res % 1_000_000_007
}
