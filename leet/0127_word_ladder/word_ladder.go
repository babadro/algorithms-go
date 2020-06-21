package _127_word_ladder

// TODO 1
func ladderLength(beginWord string, endWord string, wordList []string) int {
	//found := false
	//for _, str := range wordList {
	//	if endWord == str {}
	//	found = true
	//}
	//if !found {
	//	return 0
	//}
	//if haveOneDifferentChar(beginWord, endWord) {
	//	return 1
	//}
	res := 0
	search(wordList, beginWord, endWord, &res, 0)
	return res
}

func haveOneDifferentChar(str1, str2 string) bool {
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
	for i, w := range wordlist {
		if haveOneDifferentChar(word, w) {
			if w == endWord {
				if *res == 0 || counter < *res {
					*res = counter
				}
				return
			}
			wordlist[i] = ""
			search(wordlist, w, endWord, res, counter+1)
			wordlist[i] = w
		}
	}
}
