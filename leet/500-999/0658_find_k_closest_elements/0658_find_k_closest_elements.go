package _658_find_k_closest_elements

import "math"

// passed
func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) == 1 {
		return arr
	}

	idx := binarySearch(arr, x)
	if arr[idx] > x {
		diff1 := math.MaxInt64
		if idx-1 >= 0 {
			diff1 = abs(x - arr[idx-1])
		}

		diff2 := abs(x - arr[idx])

		if diff1 < diff2 || diff1 == diff2 {
			idx = idx - 1
		}
	}

	k--
	l, r := idx, idx
	for k > 0 {
		diff1, diff2 := math.MaxInt64, math.MaxInt64
		if l-1 >= 0 {
			diff1 = abs(x - arr[l-1])
		}

		if r+1 < len(arr) {
			diff2 = abs(x - arr[r+1])
		}

		if diff1 < diff2 || diff1 == diff2 {
			l--
		} else {
			r++
		}

		k--
	}

	return arr[l : r+1]
}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := l + (r-l)/2
		num := arr[m]
		if num > target {
			r = m
		} else if num < target {
			l = m + 1
		} else {
			return m
		}
	}

	return l
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}
