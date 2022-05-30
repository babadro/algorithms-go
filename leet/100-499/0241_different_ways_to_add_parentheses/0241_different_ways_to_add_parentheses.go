package _241_different_ways_to_add_parentheses

import "strconv"

// tptl. passed
func diffWaysToCompute(expression string) []int {
	return rec([]byte(expression), make(map[string][]int))
}

func rec(b []byte, m map[string][]int) []int {
	if v, ok := m[string(b)]; ok {
		return v
	}

	var results []int
	for i, chr := range b {
		if chr == '+' || chr == '-' || chr == '*' {
			leftParts := rec(b[0:i], m)
			rightParts := rec(b[i+1:], m)

			var res int
			for _, part1 := range leftParts {
				for _, part2 := range rightParts {
					switch chr {
					case '+':
						res = part1 + part2
					case '-':
						res = part1 - part2
					case '*':
						res = part1 * part2
					}

					results = append(results, res)
				}
			}
		}
	}

	if len(results) == 0 {
		num, _ := strconv.Atoi(string(b))
		results = []int{num}
	}

	m[string(b)] = results

	return results
}
