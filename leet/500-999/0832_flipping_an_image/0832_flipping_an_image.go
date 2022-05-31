package _832_flipping_an_image

// tptl. passed
func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		for i := 0; i < (len(row)+1)/2; i++ {
			left, right := i, len(row)-1-i
			row[left], row[right] = row[right]^1, row[left]^1
		}
	}

	return image
}
