package _365_how_many_numbers_are_smaller_than_the_curr_num

import "testing"

func TestSmallerNumbersThanCurrent(t *testing.T) {
	t.Log(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	t.Log(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))
	t.Log(smallerNumbersThanCurrent([]int{7, 7, 7, 7}))
}
