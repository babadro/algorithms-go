package _033_search_in_rotated_sorted_array

func search2(nums []int, target int) int {
	n := len(nums)
	start1, start2 := 0, n
	start2 = findStart2(nums)

	if target < nums[0] {
		return binarySearch2(nums, start2, n, target)
	}

	return binarySearch2(nums, start1, start2, target)
}

func findStart2(nums []int) int {
	l, r, res := 0, len(nums), -1
	first, last := nums[0], nums[len(nums)-1]

	for l < r {
		mid := (l + r) / 2
		curr := nums[mid]
		if curr > last {
			l = mid + 1
		} else if curr < first {
			res, r = mid, mid
		} else {
			return len(nums)
		}
	}

	return res
}

func binarySearch2(nums []int, left, right, target int) int {
	for left < right {
		mid := (left + right) / 2
		curr := nums[mid]
		if target > curr {
			left = mid + 1
		} else if target < curr {
			right = mid
		} else {
			return mid
		}
	}

	return -1
}
