package _986_interval_list_intersection

// tptl. passed.
func intervalIntersection2(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		a, b := firstList[i], secondList[j]
		if (a[0] >= b[0] && a[0] <= b[1]) || (b[0] >= a[0] && b[0] <= a[1]) {
			res = append(res, []int{max(a[0], b[0]), min(a[1], b[1])})
		}

		if a[1] < b[1] {
			i++
		} else {
			j++
		}
	}

	return res
}

// passed
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		a, b := firstList[i], secondList[j]
		if in := intersect(a, b); in != nil {
			res = append(res, in)
		}

		if a[1] < b[1] {
			i++
		} else if b[1] < a[1] {
			j++
		} else {
			i++
			j++
		}
	}

	return res
}

func intersect(a, b []int) []int {
	if a[0] > b[0] {
		a, b = b, a
	}

	if a[1] < b[0] {
		return nil
	}

	return []int{b[0], min(a[1], b[1])}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
