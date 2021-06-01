package _1881_maximum_value_after_insertion

import "strconv"

// passed. not difficult
func maxValue(n string, x int) string {
	i := 0
	negative := n[0] == '-'
	if negative {
		i = 1
	}

	for ; i < len(n); i++ {
		char := n[i]
		digit := int(char - '0')

		if (negative && x < digit) || (!negative && x > digit) {
			return n[:i] + strconv.Itoa(x) + n[i:]
		}

	}

	return n + strconv.Itoa(x)
}
