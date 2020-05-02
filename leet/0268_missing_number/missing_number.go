package _268_missing_number

func missingNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	expectedSum := (len(nums) * (len(nums) + 1)) / 2
	return expectedSum - sum
}
