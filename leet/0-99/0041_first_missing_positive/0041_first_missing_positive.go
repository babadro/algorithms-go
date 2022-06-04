package _041_first_missing_positive

func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] != i+1 {
			tmp := nums[i]
			nums[i] = nums[nums[i]-1]
			nums[tmp-1] = tmp
		}
	}

	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}

	return -1
}

for i := 0; i < len(nums); i++ {
for nums[i] != i+1 {
tmp := nums[i]
nums[i] = nums[nums[i]-1]
nums[tmp-1] = tmp
}
}