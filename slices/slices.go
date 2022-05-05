package slices

import "sort"

func IntSlicesAreEqual(slice1, slice2 []int) bool {
	length := len(slice1)
	if length != len(slice2) {
		return false
	}
	for i := 0; i < length; i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func SliceOfIntSlicesAreEqual(slice1, slice2 [][]int) bool {
	length := len(slice1)
	if length != len(slice2) {
		return false
	}
	for i := 0; i < length; i++ {
		if !IntSlicesAreEqual(slice1[i], slice2[i]) {
			return false
		}
	}
	return true
}

func SortEach(arr [][]int) {
	for i := range arr {
		sort.Ints(arr[i])
	}
}

func Less(arr1, arr2 []int) bool {
	for i := 0; i < len(arr1) && i < len(arr2); i++ {
		if arr1[i] != arr2[i] {
			return arr1[i] < arr2[i]
		}
	}

	return false
}
