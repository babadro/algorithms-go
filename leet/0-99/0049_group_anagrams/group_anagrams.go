package _049_group_anagrams

// passed. skipped easy
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]byte][]string)
	for _, str := range strs {
		key := [26]byte{}
		for i := range str {
			char := str[i]
			key[char-'a']++
		}

		m[key] = append(m[key], str)
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}

	return res
}
