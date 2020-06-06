package _075_sort_colors

import "testing"

func TestSetColors(t *testing.T) {
	input := []int{0, 1, 0, 1, 0, 1, 0, 2, 2, 0, 0, 0}
	sortColors(input)
	t.Log(input)
	input2 := []int{2, 2, 1, 1, 0, 0}
	sortColors(input2)
	t.Log(input2)
	input3 := []int{1, 2, 0, 2, 1, 0, 1, 1, 1, 0, 1}
	sortColors(input3)
	t.Log(input3)
}
