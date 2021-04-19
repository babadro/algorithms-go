package _1539_kth_missing_positive_number

// best solution. tptl. passed
func findKthPositive2(arr []int, k int) int {
	last := 0
	for _, num := range arr {
		if num-last > k {
			return last + k
		}

		k -= num - last - 1
		last = num
	}

	return last + k
}

// passed.
func findKthPositive(arr []int, k int) int {
	last, counter := 0, 0
	for _, num := range arr {
		count := num - last - 1
		if counter+count >= k {
			return last - counter + k
		}

		counter += count

		last = num
	}

	return k - counter + arr[len(arr)-1]
}
