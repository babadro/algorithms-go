package __single_number

func findSingleNumber(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res ^= arr[i]
	}

	return res
}
