package _418_sentence_screen_fitting

func wordsTyping(sentence []string, rows int, cols int) int {
	var res int

	col, row, whiteSpace := 0, 0, false
	for {
		for i := 0; i < len(sentence); {
			if row == rows {
				return res
			}

			if col == cols {
				col, row = 0, row+1

				continue
			}

			s := sentence[i]
			lenght := len(s)
			if whiteSpace {
				lenght += 1
			}

			if col+lenght > cols {
				col, row = 0, row+1

				whiteSpace = false

				continue
			}

			col += lenght

			whiteSpace = true

			i++
		}

		res++
	}

	return res
}
