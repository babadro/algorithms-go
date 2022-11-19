package _438_find_all_anagrams_in_a_string

// based on 567
func findAnagrams2(s string, p string) []int {
	freq := make(map[byte]int)
	for i := range p {
		freq[p[i]]++
	}

	left, matched, res := 0, 0, make([]int, 0)
	for right := range s {
		rightChar := s[right]
		if _, ok := freq[rightChar]; ok {
			freq[rightChar]--

			if freq[rightChar] == 0 {
				matched++
			}
		}

		left = right - len(p) + 1
		if matched == len(freq) {
			res = append(res, left)
		}

		if left >= 0 {
			leftChar := s[left]
			if count, ok := freq[leftChar]; ok {
				if count == 0 {
					matched--
				}

				freq[leftChar]++
			}
		}
	}

	return res
}

// tptl. passed
func findAnagrams(s string, p string) []int {
	freq := make(map[byte]int)
	for i := range p {
		freq[p[i]]++
	}

	counter := len(p)
	l := 0
	var res []int
	for r := range s {
		rightCh := s[r]

		if count, ok := freq[rightCh]; ok {
			freq[rightCh]--
			if count > 0 {
				counter--
			}
		}

		for ; counter == 0; l++ {
			if r-l+1 == len(p) {
				res = append(res, l)
			}

			leftCh := s[l]
			if count, ok := freq[leftCh]; ok {
				freq[leftCh]++
				if count >= 0 {
					counter++
				}
			}
		}
	}

	return res
}
