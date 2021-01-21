package sorting

import "testing"

func TestQuickSortRecursive(t *testing.T) {
	f := func(arr []int) {
		QuickSortRecursive(arr, 0, len(arr)-1)
	}

	testFunc(t, f)
}
