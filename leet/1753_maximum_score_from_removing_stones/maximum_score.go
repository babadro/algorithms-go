package _753_maximum_score_from_removing_stones

// passed. tptl todo 2 refac
func maximumScore(a int, b int, c int) int {
	arr := [3]int{a, b, c}
	counter := 0

Loop:
	for {
		if (arr[0] == 0 && arr[1] == 0) || (arr[1] == 0 && arr[2] == 0) || (arr[0] == 0 && arr[2] == 0) {
			break
		}
		min, minI, max, maxI := 0, -1, 0, -1
		if arr[0] == arr[1] && arr[1] == arr[2] {
			arr[0]--
			arr[1]--
			counter++
			continue Loop
		}
		for i, num := range arr {
			if num == 0 {
				if i == 0 {
					arr[1]--
					arr[2]--
					counter++
					continue Loop
				} else if i == 1 {
					arr[0]--
					arr[2]--
					counter++
					continue Loop
				}

				arr[0]--
				arr[1]--
				counter++
				continue Loop

			}

			if minI == -1 || num < min {
				minI, min = i, num
			}
			if maxI == -1 || num > max {
				maxI, max = i, num
			}
		}

		arr[minI]--
		arr[maxI]--

		counter++
	}

	return counter
}
