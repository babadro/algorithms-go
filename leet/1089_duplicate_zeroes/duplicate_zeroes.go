package _1089_duplicate_zeroes

// tptl. passed best solution
func duplicateZeros2(arr []int) {
	zeroCount := 0
	for _, num := range arr {
		if num == 0 {
			zeroCount++
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if i+zeroCount < len(arr) {
				arr[i+zeroCount] = 0
			}

			zeroCount--
		}

		if i+zeroCount < len(arr) {
			arr[i+zeroCount] = arr[i]
		}
	}
}

// bruteforce. passed
func duplicateZeros(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			copy(arr[i+1:n], arr[i:n-1])
			i++
		}
	}
}
