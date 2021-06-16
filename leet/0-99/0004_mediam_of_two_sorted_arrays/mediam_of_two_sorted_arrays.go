package _004_mediam_of_two_sorted_arrays

// todo 1 O(log(m+n))
// bruteforce O(m+n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	n := n1 + n2

	if n == 0 {
		return 0
	}

	mid := n / 2
	median, prevMedian := 0, 0
	for i, j, count := 0, 0, 0; count <= mid; count++ {
		prevMedian = median
		if i < n1 && j < n2 {
			if nums1[i] > nums2[j] {
				median = nums2[j]
				j++
			} else {
				median = nums1[i]
				i++
			}
		} else if i < n1 {
			median = nums1[i]
			i++
		} else {
			median = nums2[j]
			j++
		}
	}

	if n%2 == 1 {
		return float64(median)
	}

	return (float64(median) + float64(prevMedian)) / 2
}
