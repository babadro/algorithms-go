package _1395_count_number_of_teams

// bruteforce. better, but not best
// todo 1 look for solution
func numTeams2(rating []int) int {
	n, counter := len(rating), 0
	for i := 0; i < n-2; i++ {
		first := rating[i]
		for j := i + 1; j < n-1; j++ {
			second := rating[j]
			for k := j + 1; k < n; k++ {
				third := rating[k]
				if second < first {
					if third < second {
						counter++
					}
				} else if third > second {
					counter++
				}
			}
		}
	}

	return counter
}
