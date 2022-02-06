package _2147_numbers_of_ways_to_divide_long_corridor

// tptl passed. medium
func numberOfWays(corridor string) int {
	seats, plants, res := 0, 0, 1
	for _, ch := range corridor {
		if ch == 'S' {
			seats++

			if seats%2 == 1 && plants > 0 {
				res *= plants + 1
				res %= 1_000_000_007

				plants = 0
			}
		} else if seats > 0 && seats%2 != 1 {
			plants++
		}
	}

	if seats == 0 || seats%2 == 1 {
		return 0
	}

	return res
}
