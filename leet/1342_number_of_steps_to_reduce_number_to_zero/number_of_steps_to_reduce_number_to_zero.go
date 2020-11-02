package _1342_number_of_steps_to_reduce_number_to_zero

// tptl. passed. easy. easy to remember.
func numberOfSteps(num int) int {
	counter := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		counter++
	}

	return counter
}

// best solution. bit manipulation.
func numberOfSteps2(num int) int {
	counter := 0

	for num > 0 {
		counter += num&1 + 1
		num = num >> 1
	}

	if counter == 0 {
		return 0
	}

	return counter - 1
}
