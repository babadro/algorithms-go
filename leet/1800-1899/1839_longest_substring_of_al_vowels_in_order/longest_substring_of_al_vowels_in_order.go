package _839_longest_substring_of_al_vowels_in_order

import "github.com/babadro/algorithms-go/utils"

const vowels = "aeiou"

// todo 1 look for this solution:
// https://leetcode.com/problems/longest-substring-of-all-vowels-in-order/discuss/1175164/C%2B%2B-O(n)
func longestBeautifulSubstring(word string) int {
	max := 0
	n, l := len(word), len(vowels)

	flag, j, counter := false, 0, 0
	for i := 0; i <= n; {
		if !flag && i < n && word[i] == 'a' {
			flag = true
		}

		if flag {
			if j == l {
				max = utils.Max(max, counter)
				flag, j, counter = false, 0, 0

				i++
				continue
			}

			if i < n && word[i] == vowels[j] {
				counter++
				i++
			} else if word[i-1] == vowels[j] {
				j++
			} else {
				flag, j, counter = false, 0, 0
			}
		} else {
			i++
		}
	}

	return max
}
