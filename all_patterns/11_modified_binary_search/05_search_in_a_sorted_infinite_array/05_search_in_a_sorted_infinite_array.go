package _5_search_in_a_sorted_infinite_array

import "math"

// leetcode (premium) https://leetcode.com/problems/search-in-a-sorted-array-of-unknown-size/
// Given an infinite sorted array (or an array with unknown size), find if a given number ‘key’ is present in the array. Write a function to return the index of the ‘key’ if it is present in the array, otherwise return -1.
// Since it is not possible to define an array with infinite (unknown) size, you will be provided with an interface ArrayReader to read elements of the array.
// ArrayReader.get(index) will return the number at index; if the array’s size is smaller than the index, it will return Integer.MAX_VALUE.
type arrayReader struct {
	arr []int
}

func (a *arrayReader) Get(idx int) int {
	if idx >= len(a.arr) {
		return math.MaxInt64
	}

	return a.arr[idx]
}

// tptl
func search(reader arrayReader, key int) int {
	start, end := 0, 1
	for reader.Get(end) < key {
		newStart := end + 1
		end += (end - start + 1) * 2
		start = newStart
	}

	for start <= end {
		m := start + (end-start)/2
		if reader.Get(m) > key {
			end = m - 1
		} else if reader.Get(m) < key {
			start = m + 1
		} else {
			return m
		}
	}

	return -1
}
