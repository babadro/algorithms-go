package _2147_numbers_of_ways_to_divide_long_corridor

// todo 1 refac. tptl passed. medium
func numberOfWays(corridor string) int {
	var sum int
	for _, ch := range corridor {
		if ch == 'S' {
			sum++
		}
	}

	if sum%2 != 0 || sum == 0 {
		return 0
	}

	if sum == 2 {
		return 1
	}

	count, start, finish := 0, -1, -1
	for i, ch := range corridor {
		if ch == 'S' {
			count++
		}

		if start == -1 && count == 2 {
			start = i + 1
		} else if finish == -1 && count == sum-1 {
			finish = i - 1

			break
		}
	}

	seats, plants, res := 0, 0, 1
	for i := start; i <= finish+1; i++ {
		ch := corridor[i]
		if ch == 'S' {
			seats++

			if seats == 1 {
				if plants > 0 {
					res *= plants + 1
					res %= 1_000_000_007

					plants = 0
				}
			} else { //seats == 2
				seats, plants = 0, 0
			}
		} else {
			plants++
		}
	}

	return res
}
