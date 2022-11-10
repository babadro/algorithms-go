package _424_longest_repeating_character_replacement

// tptl. passed. medium (hard for me)
func characterReplacement(s string, k int) int {
	freq := [26]int{}
	left, mostFreq, res := 0, 0, 0
	for right := range s {
		rightChar := s[right] - 'A'
		freq[rightChar]++
		mostFreq = max(mostFreq, freq[rightChar])

		if right-left+1-mostFreq > k {
			freq[s[left]-'A']--
			left++
		}

		res = max(res, right-left+1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
