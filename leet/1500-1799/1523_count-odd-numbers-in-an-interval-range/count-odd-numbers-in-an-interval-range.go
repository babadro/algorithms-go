package _1523_count_odd_numbers_in_an_interval_range

// tptl. passed. easy
func countOdds(low int, high int) int {
	count := high - low + 1

	if count%2 == 0 || low%2 == 0 {
		return count / 2
	}

	return count/2 + 1
}
