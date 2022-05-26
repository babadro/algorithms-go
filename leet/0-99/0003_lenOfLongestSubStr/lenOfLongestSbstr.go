package _003_lenOfLongestSubStr

// tptl passed. best solution. sliding window without frequency map
func lengthOfLongestSubstring3(s string) int {
	chars := [256]bool{}
	left, right, maxLen := 0, 0, 0
	for left < len(s) && right < len(s) {
		rightChar := s[right]
		if !chars[rightChar] {
			chars[rightChar] = true
			right++
			maxLen = max(maxLen, right-left)
		} else {
			chars[s[left]] = false
			left++
		}
	}

	return maxLen
}

// sliding window with freq map. easy to understand
func lengthOfLongestSubstring2(s string) int {
	m := make(map[byte]int)
	left := 0
	maxLen := 0
	for i := range s {
		m[s[i]]++
		for ; left < i && len(m) < i-left+1; left++ {
			leftChar := s[left]
			m[leftChar]--
			if m[leftChar] == 0 {
				delete(m, leftChar)
			}
		}

		maxLen = max(maxLen, i-left+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
