package _344_reverse_string

func reverseString(s []byte) {
	length := len(s)
	middle := length / 2
	for i := 0; i < middle; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}
