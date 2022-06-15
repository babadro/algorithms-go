package _977_squares_of_a_sorted_array

// tptl passed
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	idx := len(res) - 1
	for i, j := 0, len(nums)-1; i <= j; idx-- {
		left, right := nums[i]*nums[i], nums[j]*nums[j]
		if left > right {
			res[idx] = left
			i++
		} else {
			res[idx] = right
			j--
		}
	}

	return res
}
