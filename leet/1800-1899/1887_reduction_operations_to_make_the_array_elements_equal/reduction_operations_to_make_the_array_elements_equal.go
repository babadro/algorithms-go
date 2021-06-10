package _1887_reduction_operations_to_make_the_array_elements_equal

import "sort"

// todo 1
func reductionOperations(nums []int) int {
	dict := make([]int, 5*10_000)
	for _, num := range nums {
		dict[num-1] = 1
	}

	counter := 0
	for i, item := range dict {
		if item == 1 {
			dict[i] = counter
			counter++
		}
	}

	res := 0
	for _, num := range nums {
		res += dict[num-1]
	}

	return res
}

// short, but slow
func reductionOperationsSort(nums []int) int {
	res, n := 0, len(nums)
	sort.Ints(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i-1] != nums[i] {
			res += n - i
		}
	}

	return res
}
