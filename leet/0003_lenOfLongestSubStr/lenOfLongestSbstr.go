package _003_lenOfLongestSubStr

// SlidingWindow
func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	n := len(runes)
	set := make(map[int32]bool)
	ans, i, j := 0, 0, 0
	for i < n && j < n {
		if !set[runes[j]] {
			set[runes[j]] = true
			j++
			if res := j - i; res > ans {
				ans = res
			}
		} else {
			set[runes[i]] = false
			i++
		}
	}
	return ans
}

func lengthOfLongestSubstringOptimizedWindow(s string) int {
	runes := []rune(s)
	n := len(runes)
	ans := 0
	dict := make(map[int32]int)
	for j, i := 0, 0; j < n; j++ {
		if val, ok := dict[runes[j]]; ok {
			if val > i {
				i = val
			}
		}
		if res := j - i + 1; res > ans {
			ans = res
		}
		dict[runes[j]] = j + 1
	}
	return ans
}

func lengthOfLongestSubstringNaive(s string) int {
	myMap := make(map[int32]bool)
	counter := 0
	max := 0
	runes := []rune(s)
	i := 0
	last := 0
	for i < len(runes) {
		letter := runes[i]
		if _, ok := myMap[letter]; ok {
			if max < counter {
				max = counter
			}
			counter = 0
			for k := range myMap {
				delete(myMap, k)
			}
			last++
			i = last
		} else {
			myMap[letter] = true
			counter++
			i++
		}
	}
	if counter > max {
		return counter
	}
	return max
}
