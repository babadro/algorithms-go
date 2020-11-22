package _720_longest_word_in_dictionary

import "sort"

// best solution
// https://leetcode.com/problems/longest-word-in-dictionary/discuss/894746/Go-Solution-using-trie
type TrieNode struct {
	word string
	next [26]*TrieNode
}

func longestWord2(A []string) string {
	root := &TrieNode{}
	sort.Slice(A, func(i, j int) bool {
		if len(A[i]) == len(A[j]) {
			return A[i] < A[j]
		}

		return len(A[i]) < len(A[j])
	})

	if len(A) == 0 || len(A[0]) != 1 {
		return ""
	}

	curMax := A[0]
	root.next[A[0][0]-'a'] = &TrieNode{word: A[0]}

	for i := 1; i < len(A); i++ {
		node := root

		for j := 0; j < len(A[i])-1; j++ {
			node = node.next[A[i][j]-'a']
			if node == nil {
				break
			}
		}

		if node != nil {
			node.next[A[i][len(A[i])-1]-'a'] = &TrieNode{word: A[i]}

			if len(A[i]) > len(curMax) {
				curMax = A[i]
			}
		}
	}

	return curMax
}
