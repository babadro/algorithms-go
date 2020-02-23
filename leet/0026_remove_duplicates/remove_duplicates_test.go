package _026_remove_duplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{1, 2, 3}
	length := removeDuplicates(arr)
	t.Log(length, arr)
}
