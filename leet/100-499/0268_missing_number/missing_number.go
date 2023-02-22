package _268_missing_number

// passed. tptl
func missingNumber(nums []int) int {
	res, i := 0, 0
	for ; i < len(nums); i++ {
		res += i - nums[i]
	}

	return res + i
}

func missingNumber2(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	expectedSum := (len(nums) * (len(nums) + 1)) / 2

	return expectedSum - sum
}

// cyclic sort
func missingNumber3(nums []int) int {
	for i := range nums {
		for nums[i] < len(nums) && nums[i] != i {
			j := nums[i]
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	for i := range nums {
		if nums[i] != i {
			return i
		}
	}

	return len(nums)
}
