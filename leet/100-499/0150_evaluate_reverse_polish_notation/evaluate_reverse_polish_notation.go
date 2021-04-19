package _150_evaluate_reverse_polish_notation

import "strconv"

// 100% 79%
func evalRPN(tokens []string) int {
	var stack []int
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			res := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		case "-":
			res := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		case "/":
			res := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		case "*":
			res := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		default:
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		}
	}

	return stack[0]
}
