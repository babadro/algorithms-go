package _215_kth_largest_element_in_array

// https://leetcode.com/problems/kth-largest-element-in-an-array/discuss/538760/Go-quick-select
// TODO 2 remember
func findKthLargest2(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	pivot := nums[0]
	left, right := 1, 1
	for ; right < len(nums); right++ {
		if nums[right] < pivot {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
	}
	nums[0], nums[left-1] = nums[left-1], nums[0]
	if left == len(nums)-k+1 {
		return pivot
	} else if left < len(nums)-k+1 {
		return findKthLargest2(nums[left:], k)
	} else {
		return findKthLargest2(nums[:left-1], left+k-len(nums)-1)
	}
}

/*
// doesn't work
func findKthLargest3(nums []int, k int) int {
	k = len(nums) - k
	lo, hi := 0, len(nums)-1
	for lo < hi {
		j := partition(nums, lo, hi)
		if j < k {
			lo = j + 1
		} else if j > k {
			hi = j - 1
		} else {
			break
		}
	}

	return nums[k]
}

func partition(a []int, lo, hi int) int {
	i, j := lo, hi+1
	for {
		for i < hi && (a[i+1] < a[lo]) {
			i++
		}
		for j > lo && (a[lo] < a[j-1]) {
			j--
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}

*/
