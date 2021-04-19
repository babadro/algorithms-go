package _769_minimum_number_of_operations_to_move_all_balls_to_each_box

// passed. todo 2 find more optimal solution
func minOperations(boxes string) []int {
	var idxes []int
	for i, char := range boxes {
		if char == '1' {
			idxes = append(idxes, i)
		}
	}

	res := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		counter := 0
		for _, idx := range idxes {
			if idx == i {
				continue
			}

			if idx > i {
				counter += idx - i
			} else {
				counter += i - idx
			}
		}

		res[i] = counter
	}

	return res
}
