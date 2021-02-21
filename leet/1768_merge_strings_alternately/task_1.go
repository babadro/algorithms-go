package _768_merge_strings_alternately

// passed. easy
func mergeAlternately(word1 string, word2 string) string {
	var res []byte
	w1, w2 := []byte(word1), []byte(word2)

	for i := 0; len(w1) > 0 || len(w2) > 0; i++ {
		if i%2 == 0 {
			if len(w1) > 0 {
				res = append(res, w1[0])
				w1 = w1[1:]
			} else {
				res = append(res, w2[0])
				w2 = w2[1:]
			}
		} else {
			if len(w2) > 0 {
				res = append(res, w2[0])
				w2 = w2[1:]
			} else {
				res = append(res, w2[0])
				w1 = w1[1:]
			}
		}
	}

	return string(res)
}
