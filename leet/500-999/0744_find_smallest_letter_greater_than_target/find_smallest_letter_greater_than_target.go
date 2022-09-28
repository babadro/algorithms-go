package _744_find_smallest_letter_greater_than_target

import "sort"

// passed. tptl
func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	for l < r {
		m := l + (r-l)/2
		if letters[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}

	if letters[r] <= target {
		return letters[0]
	}

	return letters[r]
}

// using sort.Search
func nextGreatestLetter2(letters []byte, target byte) byte {
	n := len(letters)
	if target >= letters[n-1] {
		return letters[0]
	}

	idx := sort.Search(n, func(i int) bool {
		return letters[i] > target
	})

	return letters[idx]
}
