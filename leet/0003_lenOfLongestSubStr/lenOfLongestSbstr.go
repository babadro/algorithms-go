package _003_lenOfLongestSubStr

func lengthOfLongestSubstring(s string) int {
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
