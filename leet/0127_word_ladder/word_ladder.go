package _127_word_ladder

import "log"

var iteration int

// look at solution. This can't solve long input - time limit exceed
func ladderLengthOld(beginWord string, endWord string, wordList []string) int {
	res := 0
	search(wordList, beginWord, endWord, &res, 0)
	if res == 0 {
		return 0
	}
	return res + 1
}

func haveOneDifferentChar(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	diffFound := false
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			if diffFound {
				return false
			}
			diffFound = true
		}
	}
	return diffFound
}

func search(wordlist []string, word, endWord string, res *int, counter int) {
	if word == "" {
		return
	}
	iteration++
	if iteration >= 10000000 {
		log.Fatal("Time limit exceed")
	}
	counter++
	for i, w := range wordlist {
		if haveOneDifferentChar(word, w) {
			if w == endWord {
				if *res == 0 || counter < *res {
					*res = counter
				}
				return
			}
			if *res != 0 && counter >= *res {
				return
			}
			wordlist[i] = ""
			search(wordlist, w, endWord, res, counter)
			wordlist[i] = w
		}
	}
}
