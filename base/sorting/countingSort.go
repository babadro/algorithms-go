package sorting

// tptl
func Counting(arr []byte) {
	n := len(arr)

	output := make([]byte, n)

	count := make([]int, 256)

	for _, b := range arr {
		count[b]++
	}

	for i := 1; i < 256; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		output[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}
