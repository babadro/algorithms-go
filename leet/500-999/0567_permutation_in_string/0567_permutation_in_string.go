package _567_permutation_in_string

// passed. medium. best solution (universal), without nested loop
func checkInclusion3(s1 string, s2 string) bool {
	matched := 0
	freq := make(map[byte]int)
	for i := range s1 {
		freq[s1[i]]++
	}

	for end := range s2 {
		endChar := s2[end]
		if _, ok := freq[endChar]; ok {
			freq[endChar]--

			if freq[endChar] == 0 {
				matched++
			}
		}

		if matched == len(freq) {
			return true
		}

		if start := end - len(s1) + 1; start >= 0 {
			startChar := s2[start]
			if count, ok := freq[startChar]; ok {
				if count == 0 {
					matched--
				}

				freq[startChar]++
			}
		}
	}

	return false
}

// passed. easy to understand
func checkInclusion2(s1, s2 string) bool {
	dict, freq := [26]int{}, [26]int{}
	for i := range s1 {
		dict[s1[i]-'a']++
	}

	start, total := 0, 0
	for end := range s2 {
		endChar := s2[end] - 'a'
		count := dict[endChar]
		if count == 0 {
			for ; start < end; start++ {
				freq[s2[start]-'a']--
				total--
			}

			start = end + 1
			continue
		}

		freq[endChar]++
		total++

		for ; freq[endChar] > count; start++ {
			freq[s2[start]-'a']--
			total--
		}

		if total == len(s1) {
			return true
		}
	}

	return false
}

// passed. not optimal. there is a nested loop.
// it works fine only with small set of possible characters in pattern
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	perm := [26]int{}
	for i := range s1 {
		perm[s1[i]-'a']++
	}

	subStr := [26]int{}
Loop:
	for i := range s2 {
		rightCh := s2[i]
		subStr[rightCh-'a']++

		left := i - len(s1)
		if left >= 0 {
			subStr[s2[left]-'a']--
		}

		if i >= len(s1)-1 {
			for j := 0; j < 26; j++ {
				if subStr[j] != perm[j] {
					continue Loop
				}
			}

			return true
		}
	}

	return false
}
