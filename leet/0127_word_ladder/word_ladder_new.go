package _127_word_ladder

// TODO 1 need to understand
func ladderLengthCopy(beginWord string, endWord string, wordList []string) int {
	all := make(map[string]bool)
	for _, w := range wordList {
		all[w] = true
	}
	if !all[endWord] {
		return 0
	}
	bs, es := make(map[string]bool), make(map[string]bool)
	bs[beginWord] = true
	es[endWord] = true

	steps := 1
	for len(bs) > 0 && len(es) > 0 {
		if len(bs) > len(es) {
			bs, es = es, bs
		}
		for w, _ := range bs {
			delete(all, w)
		}
		next := make(map[string]bool)
		for w, _ := range bs {
			chars := []byte(w)
			for i, v := range chars {
				for j := byte('a'); j <= 'z'; j++ {
					if v == j {
						continue
					}
					chars[i] = j
					candi := string(chars)
					if es[candi] {
						return steps + 1
					}
					if !bs[candi] && all[candi] {
						next[candi] = true
					}
					chars[i] = v
				}
			}
		}
		bs = next
		steps++
	}

	return 0
}

// Not mine. Best speed solution
// 100% 85%
func ladderLength(beginWord string, endWord string, wordList []string) int {
	all := make(map[string]bool)
	// begineWord may not be in wordList while endWord must be in wordList
	for _, w := range wordList {
		all[w] = true
	}
	if !all[endWord] {
		return 0
	}
	bs, es := make(map[string]bool), make(map[string]bool)
	bs[beginWord] = true
	es[endWord] = true

	steps := 1
	for len(bs) > 0 && len(es) > 0 {
		if len(bs) > len(es) {
			bs, es = es, bs // make sure that bs is smaller for efficiency
		}
		for w, _ := range bs {
			delete(all, w)
		}
		next := make(map[string]bool)
		for w, _ := range bs {
			/*
			   If we try to change each character of w and check if it's in es or all, it takes O(all chars in bs) * O(26) * O(1)  map check
			   If we traverse wordList, to compare if there's one char diff, it's o(all chars in wordList * len(w)) (assume that all words have same length), which is slower
			*/
			chars := []byte(w)
			for i, v := range chars {
				for j := byte('a'); j <= 'z'; j++ {
					if v == j {
						continue
					}
					chars[i] = j
					candi := string(chars)
					if es[candi] {
						return steps + 1
					}
					if !bs[candi] && all[candi] {
						next[candi] = true
					}
					chars[i] = v
				}
			}
		}
		bs = next
		steps++
	}

	return 0
}

// Not mine. 58%, 57%
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	words := make(map[string]bool)
	for _, word := range wordList {
		words[word] = true
	}
	if !words[endWord] {
		return 0
	}
	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	length := 1
	for len(queue) > 0 {
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			lastWord := queue[0]
			queue = queue[1:]
			for j := 0; j < len(lastWord); j++ {
				for k := 'a'; k <= 'z'; k++ {
					newWord := lastWord[:j] + string(k) + lastWord[j+1:]
					if !words[newWord] || newWord == lastWord {
						continue
					}
					if newWord == endWord {
						return length + 1
					}
					if !visited[newWord] {
						queue = append(queue, newWord)
						visited[newWord] = true
					}
				}
			}
		}
		length++
	}

	return 0
}
