package _151_reverse_words_in_a_string

// passed. medium. #string
func reverseWords(s string) string {
	arr := make([][2]int, 0)
	start := -1
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			if start == -1 {
				continue
			}

			arr = append(arr, [2]int{start, i})
			start = -1
		} else if start == -1 {
			start = i
		}
	}

	b := make([]byte, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		pair := arr[len(arr)-1-i]
		if i > 0 {
			b = append(b, ' ')
		}

		begin, end := pair[0], pair[1]
		b = append(b, s[begin:end]...)
	}

	return string(b)
}
