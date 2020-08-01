package _131_palindrome_partitioning

// 18 % 20%
func partition(s string) [][]string {
	if len(s) == 0 {
		return nil
	}

	runes := []rune(s)
	res := make([][][]string, len(runes))

	last := len(runes) - 1

	res[last] = [][]string{{string(runes[last])}}

	for i := len(runes) - 2; i >= 0; i-- {
		for end := len(runes); end > i; end-- {
			candidate := runes[i:end]
			if isPalindrome(candidate) {
				if end == len(runes) {
					res[i] = append(res[i], [][]string{{string(candidate)}}...)
				} else {
					for _, arr := range res[end] {
						newArr := append([]string{string(candidate)}, arr...)
						res[i] = append(res[i], newArr)
					}
				}

			}
		}
	}

	return res[0]
}

func isPalindrome(r []rune) bool {
	n := len(r)
	for i := 0; i < n/2; i++ {
		if r[i] != r[n-1-i] {
			return false
		}
	}
	return true
}
