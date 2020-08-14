package _162_find_peak_element

func findPeakElementIterative(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func findPeakElementRecursive(nums []int) int {
	return search(nums, 0, len(nums)-1)
}

func search(nums []int, l, r int) int {
	if l == r {
		return l
	}
	mid := (l + r) / 2
	if nums[mid] > nums[mid+1] {
		return search(nums, l, mid)
	}
	return search(nums, mid+1, r)
}

// O(n) O(1)
func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return n - 1
}
