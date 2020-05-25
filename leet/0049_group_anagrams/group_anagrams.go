package _049_group_anagrams

// 100% 100%
func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]byte]int)
	var res [][]string
	for _, str := range strs {
		key := [26]byte{}
		for _, b := range str {
			key[b-'a']++
		}
		if id, ok := anagrams[key]; !ok {
			id = len(res)
			anagrams[key] = id
			res = append(res, []string{str})
		} else {
			res[id] = append(res[id], str)
		}
	}
	return res
}
