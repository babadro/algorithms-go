package _249_strobogrammat_number_2

// #bnsrg passed medium
func findStrobogrammatic(n int) []string {
	if n == 1 {
		return []string{"0", "1", "8"}
	}

	var res []string

	build(make([]byte, n), 0, &res)

	return res
}

func build(b []byte, i int, res *[]string) {
	if i > len(b)/2 || (i == len(b)/2 && len(b)%2 == 0) {
		for j, k := 0, len(b)-1; j < k; j, k = j+1, k-1 {
			switch b[j] {
			case '6':
				b[k] = '9'
			case '9':
				b[k] = '6'
			default:
				b[k] = b[j]
			}
		}

		*res = append(*res, string(b))

		return
	}

	for _, ch := range []byte{'0', '1', '6', '8', '9'} {
		if (i == 0 && ch == '0') ||
			(i == len(b)/2 && ((ch == '6') || ch == '9')) {
			continue
		}

		b[i] = ch

		build(b, i+1, res)
	}
}
