package _376_wiggle_subsequence

// todo 2 Follow up: Could you solve this in O(n) time?

type key struct {
	prev, cur int
	asc       bool
}

func wiggleMaxLengthTopDown(nums []int) int {
	dp := make(map[key]int)

	c1 := recTopDown(dp, nums, -1, 0, true)
	c2 := recTopDown(dp, nums, -1, 0, false)

	return max(c1, c2)
}

func recTopDown(dp map[key]int, nums []int, prev, cur int, asc bool) int {
	if cur == len(nums) {
		return 0
	}

	k := key{prev, cur, asc}
	res, ok := dp[k]
	if !ok {
		var c1, c2 int
		if match(nums, prev, cur, asc) {
			c1 = 1 + recTopDown(dp, nums, cur, cur+1, !asc)
		}

		c2 = recTopDown(dp, nums, prev, cur+1, asc)

		res = max(c1, c2)
		dp[k] = res
	}

	return res
}

func wiggleMaxLengthBruteForce(nums []int) int {
	c1 := recBruteForce(nums, -1, 0, true)
	c2 := recBruteForce(nums, -1, 0, false)

	return max(c1, c2)
}

func recBruteForce(nums []int, prev, cur int, asc bool) int {
	if cur == len(nums) {
		return 0
	}

	var c1, c2 int
	if match(nums, prev, cur, asc) {
		c1 = 1 + recBruteForce(nums, cur, cur+1, !asc)
	}

	c2 = recBruteForce(nums, prev, cur+1, asc)

	return max(c1, c2)
}

func match(nums []int, prev, cur int, asc bool) bool {
	if prev == -1 {
		return true
	}

	if asc {
		return nums[prev] < nums[cur]
	}

	return nums[prev] > nums[cur]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
