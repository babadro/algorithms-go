package _6_minimum_difference_element

// tptl
func searchMinDiffElement(arr []int, key int) int {
	if key < arr[0] {
		return arr[0]
	} else if key > arr[len(arr)-1] {
		return arr[len(arr)-1]
	}

	l, r := 0, len(arr)-1
	for l <= r {
		m := l + (r-l)/2
		if arr[m] > key {
			r = m - 1
		} else if arr[m] < key {
			l = m + 1
		} else {
			return arr[m]
		}
	}

	if diff(arr[l], key) < diff(arr[r], key) {
		return arr[l]
	}

	return arr[r]
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
