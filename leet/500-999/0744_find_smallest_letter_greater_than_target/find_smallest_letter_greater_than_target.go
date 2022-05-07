package _744_find_smallest_letter_greater_than_target

import "sort"

// passed. tptl
func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	for left <= right {
		mid := left + (right-left)/2
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left < len(letters) {
		return letters[left]
	}

	return letters[0]
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
