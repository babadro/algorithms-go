package _2172_maximum_and_sum_of_array

import (
	"math"
	"strconv"
)

// https://leetcode.com/problems/maximum-and-sum-of-array/discuss/1769139/Python.-Simple-dp-solution-without-bitmask-with-explanation
// passed. Not fast, but clear and simple
func maximumANDSum(nums []int, numSlots int) int {
	room, _ := strconv.Atoi("222222222"[:numSlots])
	memo := make(map[int]int)
	return dp(0, room, numSlots, nums, memo)
}

func dp(pos, room, numSlots int, nums []int, memo map[int]int) int {
	if pos == len(nums) {
		return 0
	}

	if memo[room] > 0 {
		return memo[room]
	}

	for slot := 1; slot <= numSlots; slot++ {
		base := int(math.Pow(10, float64(slot)-1))

		left := (room / base) % 10
		if left > 0 {
			memo[room] = max(memo[room], (nums[pos]&slot)+dp(pos+1, room-base, numSlots, nums, memo))
		}
	}

	return memo[room]
}

// https://leetcode.com/problems/maximum-and-sum-of-array/discuss/1766824/JavaC%2B%2BPython-DP-Solution
// passed. hard
// todo 2 - need to understand
func maximumANDSum2(nums []int, numSlots int) int {
	mask := int(math.Pow(3, float64(numSlots))) - 1
	memo := make([]int, mask+1)

	return dp2(len(nums)-1, mask, numSlots, memo, nums)
}

func dp2(i, mask, numSlots int, memo, nums []int) int {
	if memo[mask] > 0 {
		return memo[mask]
	}

	if i < 0 {
		return 0
	}

	for slot, bit := 1, 1; slot <= numSlots; slot, bit = slot+1, bit*3 {
		if (mask/bit)%3 > 0 {
			memo[mask] = max(memo[mask], (nums[i]&slot)+dp2(i-1, mask-bit, numSlots, memo, nums))
		}
	}

	return memo[mask]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
