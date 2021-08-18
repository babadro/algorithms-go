package _416_partition_equal_subset_sum

// passed #dp #topdown #medium
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, sum+1)
	}

	return req(nums, 0, sum/2, dp)
}

func req(nums []int, idx, amount int, dp [][]int) bool {
	if idx == len(nums) || amount < 0 {
		return false
	}

	if amount == 0 {
		return true
	}

	if val := dp[idx][amount]; val != 0 {
		return val == 1
	}

	res := req(nums, idx+1, amount, dp) || req(nums, idx+1, amount-nums[idx], dp)

	if res == true {
		dp[idx][amount] = 1
	} else {
		dp[idx][amount] = 2
	}

	return res
}

func canPartitionBottomUp(nums []int) bool {

}
