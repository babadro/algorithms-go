package _3_next_letter

// leetcode 744
func searchNextLetter(letters string, key byte) byte {
	left, right := 0, len(letters)-1
	for left <= right {
		mid := left + (right-left)/2
		if letters[mid] <= key {
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
