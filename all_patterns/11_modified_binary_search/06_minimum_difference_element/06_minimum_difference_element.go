package _6_minimum_difference_element

import "github.com/babadro/algorithms-go/utils"

// tptl
func searchMinDiffElement(arr []int, key int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := l + (r-l)/2
		if num := arr[m]; num > key {
			r = m
		} else if num < key {
			l = m + 1
		} else {
			return num
		}
	}

	if l-1 < 0 {
		return arr[l]
	}

	num1, num2 := arr[l-1], arr[l]
	if utils.Abs(num2-key) < utils.Abs(num1-key) {
		return num2
	}

	return num1
}
