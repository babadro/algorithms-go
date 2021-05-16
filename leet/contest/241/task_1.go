package _41

// passed. tptl. bruteforce. todo 1 bitmask solution https://leetcode.com/problems/sum-of-all-subset-xor-totals/discuss/1211182/Bitmask
func subsetXORSum(nums []int) int {
	res, combination := 0, make([]int, 0, len(nums))
	for n := 1; n <= len(nums); n++ {
		helper(nums, combination, 0, n, &res)
	}

	return res
}

func helper(nums, combination []int, i, n int, sum *int) {
	if n == 0 {
		*sum += xor(combination)
		return
	}

	for j := i; j <= len(nums)-n; j++ {
		helper(nums, append(combination, nums[j]), j+1, n-1, sum)
	}
}

func xor(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}

	return res
}
