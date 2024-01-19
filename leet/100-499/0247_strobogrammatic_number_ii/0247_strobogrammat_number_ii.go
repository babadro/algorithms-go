package _247_strobogrammatic_number_2

// #bnsrg passed medium
func findStrobogrammatic(n int) []string {
	if n == 1 {
		return []string{"0", "1", "8"}
	}

	var res []string

	build(make([]byte, n), 0, n-1, &res)

	return res
}

func build(b []byte, i, j int, res *[]string) {
	if i > j {
		*res = append(*res, string(b))

		return
	}

	for _, ch := range []byte{'0', '1', '6', '8', '9'} {
		if (i == 0 && ch == '0') ||
			(i == j && ((ch == '6') || ch == '9')) {
			continue
		}

		b[i] = ch

		switch ch {
		case '6':
			b[j] = '9'
		case '9':
			b[j] = '6'
		default:
			b[j] = ch
		}

		build(b, i+1, j-1, res)
	}
}
