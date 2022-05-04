package _4_fruits_into_baskets

import "github.com/babadro/algorithms-go/utils"

// leetcode 904
func findMaxNumFruits(arr string) int {
	var fruit1, fruit2 byte
	var basket1, basket2 int

	maxNum := 0
	for i := 0; i < len(arr); i++ {
		fruit := arr[i]
		if fruit != fruit1 && fruit != fruit2 {
			fruit1, basket1 = fruit2, basket2
			fruit2, basket2 = fruit, 1
		} else if fruit == fruit1 {
			basket1++
		} else {
			basket2++
		}

		maxNum = utils.Max(maxNum, basket1+basket2)
	}

	return maxNum
}
