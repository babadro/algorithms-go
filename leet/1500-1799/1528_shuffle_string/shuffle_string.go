package _528_shuffle_string

// tptl
// passed. best solution
func restoreString2(s string, indices []int) string {
	runes := []rune(s)
	res := make([]rune, len(runes))
	for i := range indices {
		char := runes[i]
		res[indices[i]] = char
	}

	return string(res)
}

// too complicated. passed
func restoreString(s string, indices []int) string {
	i := 0
	arr := []rune(s)
	for i < len(arr) {
		if indices[i] == i {
			i++
			continue
		}
		arr[i], arr[indices[i]] = arr[indices[i]], arr[i]
		indices[i], indices[indices[i]] = indices[indices[i]], indices[i]
	}

	return string(arr)
}

// not mine. doesn't work with non utf symbols
func restoreString3(s string, indices []int) string {
	result := make([]byte, len(indices))
	for i := 0; i < len(indices); i++ {
		index := indices[i]
		char := s[i]
		result[index] = char
	}

	return string(result)
}
