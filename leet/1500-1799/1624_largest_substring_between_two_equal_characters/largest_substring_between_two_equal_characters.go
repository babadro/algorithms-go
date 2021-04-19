package _1624_largest_substring_between_two_equal_characters

// tptl passed
func maxLengthBetweenEqualCharacters(s string) int {
	idxes := [26][2]int{}

	max := -1
	for i := range s {
		char := s[i] - 'a'
		idx, ok := idxes[char][0], idxes[char][1]
		if ok == 0 {
			idxes[char] = [2]int{i, 1}
		} else if n := i - idx - 1; n > max {
			max = n
		}
	}

	return max
}
