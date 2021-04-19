package _1346_check_if_n_and_its_double_exist

// passed. tptl. best solution
func checkIfExist2(arr []int) bool {
	m := make(map[int]bool)

	for _, num := range arr {
		if m[num*2] || (num%2 == 0 && m[num/2]) {
			return true
		}

		m[num] = true
	}

	return false
}

// passed
func checkIfExist(arr []int) bool {
	m := make(map[int]bool, len(arr))
	for _, num := range arr {
		if num == 0 && m[0] {
			return true
		}
		m[num] = true
	}

	for num := range m {
		if num != 0 && m[num*2] {
			return true
		}
	}

	return false
}
