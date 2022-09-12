package _438_find_all_anagrams_in_a_string

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
