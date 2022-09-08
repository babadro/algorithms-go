package _658_find_k_closest_elements

import "math"

// tptl. passed. beautiful, simple and short.
// Fast when k ~= len(arr), slow, when k significantly less, than len(arr)
func findClosestElements2(arr []int, k, x int) []int {
	left, right := 0, len(arr)-1
	for right-left+1 > k {
		if abs(arr[left]-x) > abs(arr[right]-x) {
			left++
		} else {
			right--
		}
	}

	return arr[left : right+1]
}

// passed, but long. Should be fast on big arr and small k, but slow when len(arr) ~= k
// O(log(n)) + O(k), where n = len(arr)
func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) == 1 {
		return arr
	}

	idx := binarySearch(arr, x)
	if arr[idx] > x {
		idx = findClosest(arr, x, idx-1, idx)
	}

	l, r := idx, idx
	for k = k - 1; k > 0; k-- {
		closestIDx := findClosest(arr, x, l-1, r+1)
		if closestIDx == l-1 {
			l--
		} else {
			r++
		}
	}

	return arr[l : r+1]
}

func findClosest(arr []int, target, l, r int) int {
	diff1, diff2 := math.MaxInt64, math.MaxInt64
	if l >= 0 {
		diff1 = abs(target - arr[l])
	}

	if r < len(arr) {
		diff2 = abs(target - arr[r])
	}

	if diff1 < diff2 || diff1 == diff2 {
		return l
	}

	return r
}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := l + (r-l)/2
		num := arr[m]
		if num >= target {
			r = m
		} else if num < target {
			l = m + 1
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
