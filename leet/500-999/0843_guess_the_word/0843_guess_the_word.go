package _843_guess_the_word

type Master interface {
	Guess(word string) int
}

// #bnsrg
func findSecretWord(words []string, master *Master) {
	l := len(words)
	k := 0
	answ := 0
	for answ != 6 {
		w := words[k]

		answ = (*master).Guess(words[k])
		// remove this if block?
		if answ == -1 {
			break
		}

		nk := 0
		for i := 0; i < l; i++ {
			if words[i] == "" {
				continue
			}

			if answ == match(w, words[i]) {
				nk = i
			} else {
				words[i] = ""
			}
		}

		k = nk
	}
}

func match(s1 string, s2 string) int {
	c := 0
	for i := range s1 {
		if s1[i] == s2[i] {
			c++
		}
	}

	return c
}
