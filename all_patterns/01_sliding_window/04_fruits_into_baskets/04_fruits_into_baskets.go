package _4_fruits_into_baskets

import "github.com/babadro/algorithms-go/utils"

// leetcode 904 todo 2 there is a more effective solution without sliding window
func findMaxNumFruits(arr string) int {
	start, maxLen := 0, 0
	frequency := make(map[byte]int, 3)
	for i := range arr {
		fruit := arr[i]
		frequency[fruit]++
		for len(frequency) > 2 {
			startFruit := arr[start]
			frequency[startFruit]--
			if frequency[startFruit] == 0 {
				delete(frequency, startFruit)
			}

			start++
		}

		maxLen = utils.Max(maxLen, i-start+1)
	}

	return maxLen
}
