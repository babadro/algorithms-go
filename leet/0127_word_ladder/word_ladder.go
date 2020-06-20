package _127_word_ladder

// TODO 1
func ladderLength(beginWord string, endWord string, wordList []string) int {
	arr := make([][]int, len(wordList))

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

func search(sequenceLen *int) {

}
