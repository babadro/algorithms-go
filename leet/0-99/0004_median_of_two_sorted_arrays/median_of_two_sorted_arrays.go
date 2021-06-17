package _004_median_of_two_sorted_arrays

import (
	"github.com/babadro/algorithms-go/utils"
	"math"
)

// passed. tptl. hard. #array best solution O(log(m+n)). todo 2 need to understand
// https://leetcode.com/problems/median-of-two-sorted-arrays/discuss/2689/Solution-in-C%2B%2B-well-explained
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	size1, size2, size := len(nums1), len(nums2), len(nums1)+len(nums2)
	mid, l, r := (size-1)/2, 0, size1-1
	for l <= r {
		m1 := l + ((r - l) / 2)
		m2 := mid - m1
		if nums1[m1] > nums2[m2] {
			r = m1 - 1
		} else {
			l = m1 + 1
		}
	}

	a1 := math.MinInt64
	if r >= 0 {
		a1 = nums1[r]
	}

	a2 := math.MinInt64
	if mid-l >= 0 {
		a2 = nums2[mid-l]
	}

	a := max(a1, a2)

	if size%2 == 1 {
		return float64(a)
	}

	b1 := math.MaxInt64
	if l < size1 {
		b1 = nums1[l]
	}

	b2 := math.MaxInt64
	if mid-r < size2 {
		b2 = nums2[mid-r]
	}

	b := min(b1, b2)

	return float64(a+b) / 2
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// educative. doesn't work
func findMedianSortedArrays2(a []int, b []int) float64 {
	n, m := len(a), len(b)
	if n < m {
		n, m = m, n
		a, b = b, a
	}

	arraySize := n + m

	i, j, median := 0, 0, -1

	for startIndex, endIndex := 0, n; startIndex <= endIndex; {
		i = (startIndex + endIndex) / 2
		j = ((arraySize + 1) / 2) - i

		if i < n && j > 0 && b[j-1] > a[i] {
			startIndex = i + 1
		} else if i > 0 && j < m && b[j] < a[i-1] {
			endIndex = i - 1
		} else {
			if i == 0 {
				median = b[j-1]
			} else if j == 0 {
				median = a[i-1]
			} else {
				median = utils.Max(a[i-1], b[j-1])
			}

			break
		}
	}

	if arraySize%2 == 1 {
		return float64(median)
	}

	if i == n {
		return float64(median+b[j]) / 2
	}

	if j == m {
		return float64(median+a[i]) / 2
	}

	return float64(median+utils.Min(a[i], b[j])) / 2
}

// bruteforce O(m+n)
func findMedianSortedArraysBruteForce(nums1 []int, nums2 []int) float64 {
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
