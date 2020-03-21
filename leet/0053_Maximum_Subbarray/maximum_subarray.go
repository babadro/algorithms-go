package _053_Maximum_Subbarray

func maxSubArray(nums []int) int {
	minSum, maxSum, totalSum := 0, 0, 0
	minIdx, maxIdx := 0, 0
	signedOnly, unsignedOnly := true, true
	maxSigned := 0
	minSumFlag, maxSumFlag := true, true
	for i, n := range nums {
		if n > 0 {
			signedOnly = false
		} else if n < 0 {
			unsignedOnly = false
		}
		if signedOnly && n > maxSigned || maxSigned == 0 {
			maxSigned = n
		}
		totalSum += n
		if totalSum < minSum || minSumFlag {
			minSum = totalSum
			minIdx = i
			minSumFlag = false
		}
		if totalSum >= maxSum || maxSumFlag {
			maxSum = totalSum
			maxIdx = i
			maxSumFlag = false
		}
	}

	if signedOnly {
		return maxSigned
	} else if unsignedOnly {
		return totalSum
	}

	start := minIdx
	if nums[minIdx] < 0 {
		start++
	}
	res := 0
	for i := start; i <= maxIdx; i++ {
		res += nums[i]
	}
	return res
}
