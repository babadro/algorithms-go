package _168_excel_sheet_column_title

// tptl. not mine. need to remember
func convertToTitle(n int) string {
	var res []byte
	for n > 0 {
		res = append(res, 'A'+byte((n-1)%26))
		n = (n - 1) / 26
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return string(res)
}

func convertToString(n int) string {
	var res []byte

	for n > 0 {
		remainder := n % 10

		char := '0' + byte(remainder)

		res = append(res, char)

		n /= 10
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return string(res)
}
