package _084_largest_rectangle_in_histogram

// https://leetcode.com/problems/largest-rectangle-in-histogram/discuss/166276/Short-Java-solution
func largestRectangleArea(heights []int) int {
	idx, res, n := 0, 0, len(heights)
	var stack []int
	for idx <= n {
		last := len(stack) - 1
		if idx < n && (len(stack) == 0 || heights[idx] >= heights[stack[last]]) {
			stack = append(stack, idx)
			idx++
		} else {
			if len(stack) == 0 && idx == len(heights) {
				break
			}

			t := stack[last]
			stack = stack[:last]
			last--

			var tmp int
			if len(stack) == 0 {
				tmp = heights[t] * idx
			} else {
				tmp = heights[t] * (idx - stack[last] - 1)
			}

			res = max(res, tmp)
		}
	}

	return res
}

// todo 1
// best solution. need to understand
func largestRectangleArea2(heights []int) int {
	// Keep a stack of indices for the heights
	// E.g. [0,1,2,3] => heights[1,6,8,9]
	stack := make([]int, 0)

	// When a height is larger or equal to the height of the last index
	// in the stack, pop the stack and calculate the height of each pop.
	heights = append(heights, 0)

	maxArea := 0

	for i := range heights {
		last := len(stack) - 1
		for len(stack) > 0 && heights[stack[last]] >= heights[i] {
			stackHeight := heights[stack[last]]
			stack = stack[:last]
			last--

			var stackWidth int
			if len(stack) == 0 {
				stackWidth = i
			} else {
				stackWidth = i - stack[last] - 1
			}

			maxArea = max(maxArea, stackHeight*stackWidth)
		}

		stack = append(stack, i)
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
