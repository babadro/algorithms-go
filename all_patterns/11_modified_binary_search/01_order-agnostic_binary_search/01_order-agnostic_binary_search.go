package _1_order_agnostic_binary_search

func search(arr []int, key int) int {
	asc := arr[0] < arr[len(arr)-1]

	less := func(a, b int) bool {
		if asc {
			return a < b
		}

		return a > b
	}

	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if less(arr[mid], key) {
			left = mid + 1
		} else if arr[mid] != key {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
