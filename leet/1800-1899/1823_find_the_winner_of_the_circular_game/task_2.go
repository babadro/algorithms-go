package _1823_find_the_winner_of_the_circular_game

// passed. tptl. todo 1 read about josephus problem
func findTheWinner(n int, k int) int {
	m := make(map[int]bool)

	i := 0
	for n-len(m) > 1 {
		for j := 0; j < k; {
			i++
			if i > n {
				i = i % n
			}

			if !m[i] {
				j++
			}
		}

		m[i] = true
	}

	for num := 1; num <= n; num++ {
		if !m[num] {
			return num
		}
	}

	return -1
}
