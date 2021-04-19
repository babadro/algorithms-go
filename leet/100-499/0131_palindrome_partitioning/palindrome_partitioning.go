package _131_palindrome_partitioning

// 18 % 20%
// TODO 2 look for fast solution in discuss
func partition(s string) [][]string {
	if len(s) == 0 {
		return nil
	}

	res := make([][][]string, len(s))

	last := len(s) - 1

	res[last] = [][]string{{string(s[last])}}

	for i := len(s) - 2; i >= 0; i-- {
		for end := len(s); end > i; end-- {
			candidate := s[i:end]
			if isPalindrome(candidate) {
				if end == len(s) {
					res[i] = append(res[i], [][]string{{candidate}}...)
				} else {
					for _, arr := range res[end] {
						newArr := append([]string{candidate}, arr...)
						res[i] = append(res[i], newArr)
					}
				}

			}
		}
	}

	return res[0]
}

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
