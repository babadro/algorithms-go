package _031_next_permutation

// passed. tptl. medium (hard for me)
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}

	if i >= 0 {
		j := n - 1
		for ; j >= 0 && nums[j] <= nums[i]; j-- {
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	for k, m := i+1, n-1; k < m; k, m = k+1, m-1 {
		nums[k], nums[m] = nums[m], nums[k]
	}
}
