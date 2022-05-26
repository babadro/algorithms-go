package _424_longest_repeating_character_replacement

func characterReplacement(s string, k int) int {
	freq := [26]int{}
	left, mostFreq, res := 0, 0, 0
	for right := range s {
		rightChar := s[right]
		freq[rightChar-'A']++

		mostFreq = max(mostFreq, freq[rightChar-'A'])

		if right-left+1-mostFreq > k {
			leftChar := s[left]
			freq[leftChar-'A']--
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
