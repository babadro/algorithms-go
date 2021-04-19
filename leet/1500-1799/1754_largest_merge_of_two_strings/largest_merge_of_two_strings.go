package _754_largest_merge_of_two_strings

// passed. tptl. todo 2 refac
func largestMerge(word1 string, word2 string) string {
	b1, b2 := []byte(word1), []byte(word2)

	var merge []byte
	for len(b1) > 0 || len(b2) > 0 {
		if len(b1) > 0 && len(b2) > 0 {
			if b1[0] > b2[0] {
				merge = append(merge, b1[0])
				b1 = b1[1:]
			} else if b1[0] < b2[0] {
				merge = append(merge, b2[0])
				b2 = b2[1:]
			} else {
				i := 0
				for i = 0; i < len(b1) && i < len(b2) && b1[i] == b2[i]; i++ {
				}
				if i == len(b1) {
					merge = append(merge, b2[0])
					b2 = b2[1:]
				} else if i == len(b2) {
					merge = append(merge, b1[0])
					b1 = b1[1:]
				} else if b1[i] > b2[i] {
					merge = append(merge, b1[0])
					b1 = b1[1:]
				} else {
					merge = append(merge, b2[0])
					b2 = b2[1:]
				}
			}
		} else if len(b1) > 0 {
			merge = append(merge, b1[0])
			b1 = b1[1:]
		} else if len(b2) > 0 {
			merge = append(merge, b2[0])
			b2 = b2[1:]
		}
	}

	return string(merge)
}
