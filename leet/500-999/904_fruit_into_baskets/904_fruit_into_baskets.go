package _904_fruit_into_baskets

// todo 1
func totalFruit(fruits []int) int {
	fruit1, fruit2, basket1, basket2 := -1, -1, 0, 0

	maxNum := 0
	for i := 0; i < len(fruits); i++ {
		fruit := fruits[i]
		if fruit != fruit1 && fruit != fruit2 {
			if i > 0 {
				if fruits[i-1] == fruit2 {
					fruit1, basket1 = fruit2, basket2
				}
			}
			fruit2, basket2 = fruit, 1
		} else if fruit == fruit1 {
			basket1++
		} else {
			basket2++
		}

		maxNum = max(maxNum, basket1+basket2)
	}

	return maxNum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
