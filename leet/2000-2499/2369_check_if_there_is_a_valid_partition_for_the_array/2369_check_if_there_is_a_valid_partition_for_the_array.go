package _2369_check_if_there_is_a_valid_partition_for_the_array

// #bnsrg passed but long, slow and recursive
// todo 2 make recursive solution shorter
// todo 2 iterative solution
func validPartition(nums []int) bool {
	return rec(nums, 0, make(map[int]bool))
}

func rec(nums []int, idx int, dp map[int]bool) bool {
	if idx == len(nums) {
		return true
	}

	if res, exists := dp[idx]; exists {
		dp[idx] = res

		return res
	}

	if idx == len(nums)-1 {
		return false
	}

	idx1, idx2, idx3 := idx, idx+1, idx+2

	if idx == len(nums)-2 {
		return nums[idx1] == nums[idx2]
	}

	res := false

	if nums[idx1] == nums[idx2]-1 && nums[idx2] == nums[idx3]-1 {
		res = rec(nums, idx+3, dp)
	} else if nums[idx1] == nums[idx2] && nums[idx2] == nums[idx3] {
		res = rec(nums, idx+2, dp) || rec(nums, idx+3, dp)
	} else if nums[idx1] == nums[idx2] {
		res = rec(nums, idx+2, dp)
	}

	dp[idx] = res

	return res
}
