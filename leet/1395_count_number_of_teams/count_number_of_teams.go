package _1395_count_number_of_teams

const (
	first     = 0
	second    = 1
	thirdUp   = 2
	thirdDown = 3
)

// passed, but slow
func numTeams(rating []int) int {
	counter := 0

	helper(first, &counter, 0, rating)

	return counter
}

func helper(mode int, counter *int, prev int, arr []int) {
	for i, num := range arr {
		if mode == first {
			helper(second, counter, num, arr[i+1:])
		} else if mode == second {
			nextMode := thirdUp
			if num < prev {
				nextMode = thirdDown
			}

			helper(nextMode, counter, num, arr[i+1:])
		} else if (mode == thirdUp && num > prev) || (mode == thirdDown && num < prev) {
			*counter++
		}
	}
}
