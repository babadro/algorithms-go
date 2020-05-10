package _008_string_to_integer

import "math"

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	if i == len(str) {
		return 0
	}
	var signed bool
	if str[i] == '-' {
		signed = true
		i++
	} else if str[i] == '+' {
		i++
	}
	num := 0
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		digit := str[i] - '0'
		num *= 10
		num += int(digit)
		if signed {
			if -1*num <= math.MinInt32 {
				return math.MinInt32
			}
		} else if num >= math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	if signed {
		num *= -1
	}
	return num
}
