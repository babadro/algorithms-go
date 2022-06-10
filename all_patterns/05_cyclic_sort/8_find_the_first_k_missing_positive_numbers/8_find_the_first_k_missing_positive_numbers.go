package __find_the_first_k_missing_positive_numbers

// see very similar problem leetcode 41

// Given an unsorted array containing numbers and a number ‘k’,
// find the first ‘k’ missing positive numbers in the array.
func findNumbers(nums []int, k int) []int {
	for i := range nums {
		for nums[i] > 0 && nums[i] <= len(nums) {
			idx := nums[i] - 1

			if nums[i] == nums[idx] {
				break
			}

			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}

	var res []int
	extraNumbers := make(map[int]bool)
	for i := 0; i < len(nums) && len(res) < k; i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
			extraNumbers[nums[i]] = true
		}
	}

	for i := 1; len(res) < k; i++ {
		candidateNum := i + len(nums)
		if !extraNumbers[candidateNum] {
			res = append(res, candidateNum)
		}
	}

	return res
}
