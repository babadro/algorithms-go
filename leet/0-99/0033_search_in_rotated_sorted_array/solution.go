package _033_search_in_rotated_sorted_array

func search2(nums []int, target int) int {
	n := len(nums)
	start1, start2 := 0, n
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			start2 = i+1
			break
		}
	}

	if res := binarySearch2(nums, start1, start2, target); res != -1 {
		return res
	}

	return binarySearch2(nums, start2, n, target)
}

func findStart2(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l+r)/2
		curr := nums[mid]
		if
	}

	return len(nums)
}

func binarySearch2(nums []int, left, right, target int) int {
	for left < right {
		midIdx := (left+right)/2
		curr := nums[midIdx]
		if target > curr {
			left = midIdx + 1
		} else if target < curr {
			right = midIdx
		} else {
			return midIdx
		}
	}

	return -1
}
