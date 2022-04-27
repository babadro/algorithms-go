package _3_squaring_sorted_array

// todo 2 can I do it in place?
func makeSquares(arr []int) []int {
	n := len(arr)
	res := make([]int, n)

	for left, right, i := 0, n-1, n-1; left <= right; i-- {
		leftSquare, rightSquare := arr[left]*arr[left], arr[right]*arr[right]

		if leftSquare > rightSquare {
			res[i] = leftSquare
			left++
		} else {
			res[i] = rightSquare
			right--
		}
	}

	return res
}
