package _41

import "fmt"

// todo 1
func rearrangeSticks(n int, k int) int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	count := 0
	if visibleCount(nums) == k {
		fmt.Println(nums)
		count = 1
	}

	for {
		ok := nextPermutation(nums)
		if !ok {
			break
		}

		if visibleCount(nums) == k {
			count++
			fmt.Println(nums)
		}
	}

	return count // % 1_000_000_007
}

func visibleCount(nums []int) int {
	max, count := 0, 0
	for _, num := range nums {
		if num > max {
			count++
			max = num
		}
	}

	return count
}

func nextPermutation(nums []int) (ok bool) {
	n := len(nums)
	i := n - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}

	if i >= 0 {
		ok = true
		j := n - 1
		for ; j >= 0 && nums[j] <= nums[i]; j-- {
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	for k, m := i+1, n-1; k < m; k, m = k+1, m-1 {
		nums[k], nums[m] = nums[m], nums[k]
	}

	return ok
}
