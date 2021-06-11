package sorting

import "testing"

func TestSelectionSort(t *testing.T) {
	f := func(arr []int) {
		SelectionSort(arr)
	}

	testFunc(t, f)
}
