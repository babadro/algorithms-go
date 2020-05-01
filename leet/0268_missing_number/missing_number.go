package _268_missing_number

// TODO
func missingNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	length := len(nums) - 1
	expectedSum := (length * (length + 1)) / 2
	return expectedSum - sum
}
