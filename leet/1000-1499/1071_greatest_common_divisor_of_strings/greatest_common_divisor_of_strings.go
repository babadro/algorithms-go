package _1071_greatest_common_divisor_of_strings

// passed. tptl
func gcdOfStrings(str1 string, str2 string) string {
	divisor := gcd(len(str1), len(str2))

	if str1[:divisor] != str2[:divisor] {
		return ""
	}

	for i := range str1 {
		if str1[i] != str1[i%divisor] {
			return ""
		}
	}

	for i := range str2 {
		if str2[i] != str2[i%divisor] {
			return ""
		}
	}

	return str1[:divisor]
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
