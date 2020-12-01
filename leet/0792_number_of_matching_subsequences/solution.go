package _792_number_of_matching_subsequences

import "sort"

// todo 2. not fastest, but interesting. need to understand.
func numMatchingSubseq4(S string, words []string) int {
	pos := make([][]int, 26)
	for i := 0; i < len(S); i++ {
		pos[S[i]-'a'] = append(pos[S[i]-'a'], i)
	}

	count := 0
	for _, word := range words {
		if len(word) > len(S) {
			continue
		}

		if isSubsequence4(pos, word) {
			count++
		}
	}

	return count
}

func isSubsequence4(pos [][]int, word string) bool {
	pre := 0

	for i := 0; i < len(word); i++ {
		arr := pos[word[i]-'a']

		if arr == nil {
			return false
		}

		idx := sort.Search(len(arr), func(i int) bool { return arr[i] >= pre })
		if idx == len(arr) {
			return false
		}
		pre = arr[idx] + 1
	}

	return true
}

// passed. hard to understand
func numMatchingSubseq3(S string, words []string) int {
	ans := 0
	al := make([][]int, 26)

	for i := 0; i < 26; i++ {
		al[i] = make([]int, 0)
	}

	for i := 0; i < len(S); i++ {
		al[S[i]-'a'] = append(al[S[i]-'a'], i)
	}

Loop:
	for i := 0; i < len(words); i++ {
		key := 0
		for j := 0; j < len(words[i]); j++ {
			arr := al[words[i][j]-'a']
			ind := sort.Search(len(arr), func(i int) bool { return arr[i] >= key })
			if ind < len(arr) && arr[ind] == key {
				key++
			} else {
				if ind >= len(arr) {
					continue Loop
				}

				key = arr[ind] + 1
			}
		}

		ans++
	}

	return ans
}

// best speed. tptl. passed. easy.
func numMatchingSubseq2(S string, words []string) int {
	ans := 0

	m := make(map[string]int)
	for i := range words {
		m[words[i]]++
	}
	for sub, val := range m {
		idx := 0
		for i := 0; i < len(S) && idx < len(sub); i++ {
			if S[i] == sub[idx] {
				idx++
			}
		}
		if idx == len(sub) {
			ans += val
		}
	}

	return ans
}
