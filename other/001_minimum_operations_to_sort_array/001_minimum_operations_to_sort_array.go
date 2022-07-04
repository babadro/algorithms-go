package _01_minimum_operations_to_sort_array

import "sort"

// Sort the integers.
//Given a permutation of from 1 to n (an array), the only way you have to change it is to select any integer and put it into the end of the list.
//Return the minimal number of operations to sort the array.
// arr[] = {5, 1, 4, 3, 2}
//Return 3
//{5, 1, 4, 3, 2}  -> {5, 1, 4, 2, 3}  -> {5, 1, 2, 3, 4}-> {1, 2, 3, 4, 5}
//                           3                 4                     5
//arr[] = {2, 1, 4, 3,  5}
//Return 4
//{2, 1, 4, 3,  5} -> {1, 4, 3, 5, 2} ->{1, 4, 5, 2, 3}- > {1, 5, 2, 3, 4} -> {1, 2, 3, 4, 5}
//                              (2                   3   )              4                 5

// tptl medium (hard for me)
func minOperations(arr []int) int {
	// search the last element of increasing subsequence started from 1
	nextNum := 1
	for _, num := range arr {
		if num == nextNum {
			nextNum++
		}
	}

	// nextNum-1 is a len of such subsequence
	// {6, 1, 4, 2, 3, 5} - here this subsequence is {1,2,3}
	// nextNum after loop above will be equal 4.
	// so, the len of subsequence is 4-1 = 3:

	return len(arr) - (nextNum - 1)
}

// array of arbitrary integers (not permutation from 1 to an as in example above)
// https://leetcode.com/discuss/interview-question/789524/determine-minimum-numbers-of-moves-required-to-sort-the-array-in-ascending-order
func minOperations2(arr []int) int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)
	res := len(arr)
	j := 0
	for i := range arr {

	}

}
