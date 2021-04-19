package _1743_restore_the_array_from_adjacent_pairs

// contest. passed. todo 3 refac
func restoreArray(adjacentPairs [][]int) []int {
	m0, m1, counter := make(map[int][][]int), make(map[int][][]int), make(map[int]int)

	for _, pair := range adjacentPairs {
		m0[pair[0]] = append(m0[pair[0]], pair)
		m1[pair[1]] = append(m1[pair[1]], pair)
		counter[pair[0]]++
		counter[pair[1]]++
	}

	var start int
	startFound := false
	var finish int
	for num, count := range counter {
		if count == 1 {
			if !startFound {
				start = num
				startFound = true
				continue
			}

			finish = num
		}
	}

	res := []int{start}

	for {
		lastIdx := len(res) - 1
		lastNum := res[lastIdx]
		if lastNum == finish {
			break
		}

		var next int
		for _, pair := range m0[lastNum] {
			if lastNum != start {
				if res[lastIdx-1] != pair[1] {
					next = pair[1]
					break
				}
			} else {
				next = pair[1]
			}
		}

		for _, pair := range m1[lastNum] {
			if lastNum != start {
				if res[lastIdx-1] != pair[0] {
					next = pair[0]
					break
				}
			} else {
				next = pair[0]
			}
		}

		res = append(res, next)
	}

	return res
}

func equal(slice1, slice2 []int) bool {
	return slice1[0] == slice2[0] && slice1[1] == slice2[1]
}
