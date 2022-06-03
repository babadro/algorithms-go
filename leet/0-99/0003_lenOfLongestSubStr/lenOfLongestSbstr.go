package _003_lenOfLongestSubStr

// tptl. another one best solution
func lengthOfLongestSubstring4(s string) int {
	start, maxLen := 0, 0
	indexes := make(map[byte]int, 256)
	for end := 0; end < len(s); end++ {
		right := s[end]
		if idx, ok := indexes[right]; ok {
			start = max(start, idx+1)
		}

		indexes[right] = end
		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}

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
