package _45

// todo 1
func mergeTriplets(triplets [][]int, target []int) bool {
	t1, t2, t3 := target[0], target[1], target[2]
	for i := 0; i < len(triplets); {
		triplet := triplets[i]
		if triplet[0] > t1 || triplet[1] > t2 || triplet[2] > t3 {
			last := len(triplets) - 1
			triplets[i], triplets[last] = triplets[last], triplets[i]

			triplets = triplets[:last]
		} else {
			i++
		}
	}

	var a, b, c bool
	for _, triplet := range triplets {
		if triplet[0] == t1 {
			a = true
		}
		if triplet[1] == t2 {
			b = true
		}
		if triplet[2] == t3 {
			c = true
		}
	}

	return a && b && c
}
