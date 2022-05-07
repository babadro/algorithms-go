package _2_ceiling_of_a_number

// leetcode 35
func searchCeilingOfANumber(arr []int, key int) int {
	if key > arr[len(arr)-1] {
		return -1
	}

	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] < key {
			left = mid + 1
		} else if arr[mid] > key {
			right = mid - 1
		} else {
			return mid
		}
	}

	return left
}
