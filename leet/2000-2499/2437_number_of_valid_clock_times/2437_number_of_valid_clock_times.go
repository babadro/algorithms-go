package _2437_number_of_valid_clock_times

// #bnsrg easy passed
func countTime(time string) int {
	first, second, third, fourth := time[0], time[1], time[3], time[4]

	res := 1

	switch {
	case first == '?' && second == '?':
		res *= 24
	case first == '?':
		secondNum := second - '0'

		if secondNum > 3 {
			res *= 2
		} else {
			res *= 3
		}
	case second == '?':
		if first == '0' || first == '1' {
			res *= 10
		} else {
			res *= 4
		}
	}

	if third == '?' {
		res *= 6
	}

	if fourth == '?' {
		res *= 10
	}

	return res
}
