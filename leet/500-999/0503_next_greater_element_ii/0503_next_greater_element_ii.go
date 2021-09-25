package _503_next_greater_element_ii

// tptl passed. medium (hard for me) look solution (animation is there)
func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	var stack []int
	for i := 2*len(nums) - 1; i >= 0; i-- {
		top := len(stack) - 1
		idx := i % len(nums)
		for len(stack) != 0 && nums[stack[top]] <= nums[idx] {
			stack = stack[:top]
			top--
		}

		if len(stack) == 0 {
			res[idx] = -1
		} else {
			res[idx] = nums[stack[top]]
		}

		stack = append(stack, idx)
	}

	return res
}

func nextGreaterElementsBruteForce(nums []int) []int {
	res := make([]int, len(nums))
Loop:
	for i, num := range nums {
		for k := i + 1; ; k++ {
			if j := k % len(nums); j == i {
				break
			} else if next := nums[j]; next > num {
				res[i] = next

				continue Loop
			}
		}

		res[i] = -1
	}

	return res
}
