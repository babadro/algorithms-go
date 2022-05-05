package _904_fruit_into_baskets

// tptl todo 2 look for faster solution
func totalFruit(fruits []int) int {
	start, maxLen := 0, 0
	frequency := make(map[int]int, 3)
	for i, fruit := range fruits {
		frequency[fruit]++
		for len(frequency) > 2 {
			startFruit := fruits[start]
			frequency[startFruit]--
			if frequency[startFruit] == 0 {
				delete(frequency, startFruit)
			}

			start++
		}

		maxLen = max(maxLen, i-start+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
