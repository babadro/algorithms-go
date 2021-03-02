package _1560_most_visited_sector_in_a_circular_track

// passed. tptl. best solution
func mostVisited2(n int, rounds []int) []int {
	result := []int{}

	start := rounds[0]
	end := rounds[len(rounds)-1]

	if start < end {
		for start <= end {
			result = append(result, start)
			start++
		}
	} else if start == end {
		result = append(result, start)
	} else {
		i := 1
		for i <= end {
			result = append(result, i)
			i++
		}

		for start <= n {
			result = append(result, start)
			start++
		}
	}

	return result
}

// passed
func mostVisited(n int, rounds []int) []int {
	counters := [100]int{}
	m := len(rounds)

	sum := rounds[0] - 1
	for i := 1; i < m; {
		sector := sum % n
		if sector == rounds[i]-1 {
			i++
		}
		counters[sector]++
		sum++
	}

	max := 0
	var res []int
	for i := 0; i < n; i++ {
		counter := counters[i]
		if counter == 0 {
			continue
		}

		if counter > max {
			max = counter
			res = append(res[:0], i+1)
		} else if counter == max {
			res = append(res, i+1)
		}
	}

	return res
}
