package _053_Maximum_Subbarray

func maxSubArray(nums []int) int {
	currSum, maxSum := nums[0], nums[0]
	for _, v := range nums[1:] {
		if currSum < 0 {
			currSum = v
		} else {
			currSum += v
		}
		if currSum > maxSum {
			maxSum = currSum
		}
	}
	return maxSum
}
