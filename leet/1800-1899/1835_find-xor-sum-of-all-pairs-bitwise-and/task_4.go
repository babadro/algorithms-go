package _1835_find_xor_sum_of_all_pairs_bitwise_and

// tptl. passed. hard.
func getXORSum(arr1 []int, arr2 []int) int {
	return arrXor(arr1) & arrXor(arr2)
}

func arrXor(arr []int) int {
	res := 0
	for _, num := range arr {
		res ^= num
	}

	return res
}
