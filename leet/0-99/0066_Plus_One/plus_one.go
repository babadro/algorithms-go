package _066_Plus_One

func plusOne(digits []int) []int {
	sum := 0
	overflow := false
	for i := len(digits) - 1; i >= 0; i-- {
		if sum = digits[i] + 1; sum == 10 {
			digits[i] = 0
		} else {
			digits[i] = sum
			break
		}
		if i == 0 {
			overflow = true
		}
	}
	if overflow {
		return append([]int{1}, digits...)
	}

	return digits
}
