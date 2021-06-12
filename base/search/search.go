package search

func BinarySearch(num int, nums []int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] > num {
			right = mid - 1
		} else if nums[mid] < num {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
