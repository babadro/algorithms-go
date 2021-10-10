package _739_daily_temperatures

// passed. tptl. #medium #monostack
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int
	for i := len(temperatures) - 1; i >= 0; i-- {
		t := temperatures[i]
		last := len(stack) - 1
		for len(stack) > 0 && temperatures[stack[last]] <= t {
			stack = stack[:last]
			last--
		}

		if len(stack) > 0 {
			res[i] = stack[last] - i
		}

		stack = append(stack, i)
	}

	return res
}

func dailyTemperatures2(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i := range temperatures {
		top := len(stack) - 1
		for len(stack) > 0 && temperatures[i] > temperatures[stack[top]] {
			ans[stack[top]] = i - stack[top]

			stack = stack[:top]
			top--
		}

		stack = append(stack, i)
	}

	return ans
}
