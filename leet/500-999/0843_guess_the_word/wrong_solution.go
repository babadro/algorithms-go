package _843_guess_the_word

func findSecretWord2(words []string, master *Master) {
	m := map[string]bool{}
	for _, w := range words {
		m[w] = true
	}

	g := ""
	for {
		// pick a word and remove from map
		for w, _ := range m {
			g = w
			delete(m, g)
			break
		}
		n := (*master).Guess(g)
		// correct guess
		if n == 6 {
			return
		}

		// remove all impossible words
		for w, _ := range m {
			if d(g, w) != n {
				delete(m, w)
			}
		}
	}
}

func d(a, b string) int {
	ans := 0
	for i := 0; i < 6; i++ {
		if a[i] == b[i] {
			ans++
		}
	}
	return ans
}
