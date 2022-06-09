package __find_the_first_k_missing_positive_numbers

// see very similar problem leetcode 41

// todo solve task according conditions:
// Given an unsorted array containing numbers and a number ‘k’,
// find the first ‘k’ missing positive numbers in the array.

func findNumbers(nums []int, k int) {
	for i := range nums {
		for nums[i] > 0 && nums[i] < len(nums)
	}
}
