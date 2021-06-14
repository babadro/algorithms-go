package _899_merge_triplets_to_form_target_triplet

// tptl. passed. medium
func mergeTriplets(triplets [][]int, target []int) bool {
	t1, t2, t3 := target[0], target[1], target[2]
	var a, b, c bool
	for _, triplet := range triplets {
		if triplet[0] > t1 || triplet[1] > t2 || triplet[2] > t3 {
			continue
		}

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
