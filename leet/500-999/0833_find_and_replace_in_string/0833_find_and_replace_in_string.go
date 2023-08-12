package _833_find_and_replace_in_string

type replacement struct {
	originalStringLen int
	replacementStrIDx int
}

// bnsrg. medium passed
// todo 2 find shorter and faster solution
func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	replacements := make([]replacement, len(s))

	for k, idx := range indices {
		i, j := idx, 0
		source := sources[k]
		for ; i < len(s) && j < len(source); i, j = i+1, j+1 {
			if s[i] != source[j] {
				break
			}
		}

		if match := j == len(source); match {
			replacements[idx] = replacement{
				originalStringLen: len(source),
				replacementStrIDx: k,
			}
		}
	}

	var res []byte
	for i := 0; i < len(s); {
		replace := replacements[i]
		if replace.originalStringLen != 0 {
			res = append(res, targets[replace.replacementStrIDx]...)

			i += replace.originalStringLen
			continue
		}

		res = append(res, s[i])
		i++
	}

	return string(res)
}
