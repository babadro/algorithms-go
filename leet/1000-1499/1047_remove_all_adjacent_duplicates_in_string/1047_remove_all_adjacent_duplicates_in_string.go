package _1047_remove_all_adjacent_duplicates_in_string

// passed. easy
func removeDuplicates(s string) string {
	var res []byte
	for i := range s {
		ch := s[i]
		if len(res) > 0 && res[len(res)-1] == ch {
			res = res[:len(res)-1]

			continue
		}

		res = append(res, ch)
	}

	return string(res)
}
