package _1887_reduction_operations_to_make_the_array_elements_equal

import (
	"math"
	"sort"
)

// passed. mine. fast. medium. Optimal solution
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

// tricky
func reductionOperations2(nums []int) int {
	smallest, largest := math.MaxInt32, 0
	counter := make([]int, 5*10_000+1)
	for _, num := range nums {
		smallest = min(num, smallest)
		largest = max(num, largest)
		counter[num]++
	}

	res, prev := 0, 0
	for i := largest; i > smallest; i-- {
		if counter[i] > 0 {
			counter[i] += prev
			res += counter[i]
			prev = counter[i]
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// short, interesting, but slow
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
