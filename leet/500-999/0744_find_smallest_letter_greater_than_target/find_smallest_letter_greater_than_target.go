package _744_find_smallest_letter_greater_than_target

import "sort"

// passed. easy. tptl
func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	if target >= letters[n-1] {
		return letters[0]
	}

	idx := sort.Search(n, func(i int) bool {
		return letters[i] > target
	})

	return letters[idx]
}

// just for refresh binary search
func nextGreatestLetter2(letters []byte, target byte) byte {
	n := len(letters)
	if target >= letters[n-1] {
		return letters[0]
	}

	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return letters[l]
}
