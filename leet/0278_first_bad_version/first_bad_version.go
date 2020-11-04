package _278_first_bad_version

const badVersion = 3

// tptl. passed
func firstBadVersion(n int) int {
	left, right := 1, n
	for left <= right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func isBadVersion(ver int) bool {
	if ver >= badVersion {
		return true
	}
	return false
}
