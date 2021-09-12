package _084_largest_rectangle_in_histogram

// todo 1
// https://leetcode.com/problems/largest-rectangle-in-histogram/discuss/166276/Short-Java-solution
func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	var stack []int
	var cur int
	area := 0
	for i := 0; i <= n; i++ {
		if i == n {
			cur = 0
		} else {
			cur = heights[i]
		}

		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if cur <= heights[top] {
				break
			}

			stack = stack[:len(stack)-1]

			var tmp int
			if len(stack) == 0 {
				tmp = heights[top] * i
			} else {
				newTop := stack[len(stack)-1]
				tmp = heights[top] * (i - (newTop + 1))
			}

			if tmp > area {
				area = tmp
			}
		}

		stack = append(stack, i)
	}

	return area
}
