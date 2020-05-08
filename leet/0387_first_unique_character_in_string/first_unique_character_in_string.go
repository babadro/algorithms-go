package _387_first_unique_character_in_string

func firstUniqChar(s string) int {
	arr := [26]int{}
	for _, r := range s {
		arr[r-'a']++
	}
	for i, r := range s {
		if arr[r-'a'] == 1 {
			return i
		}
	}
	return -1
}
