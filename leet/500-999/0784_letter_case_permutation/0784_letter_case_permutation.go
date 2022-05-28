package _784_letter_case_permutation

// tptl passed
func letterCasePermutation(s string) []string {
	var res []string
	rec([]byte(s), 0, &res)

	return res
}

func rec(arr []byte, i int, res *[]string) {
	if i == len(arr) {
		*res = append(*res, string(arr))

		return
	}

	rec(arr, i+1, res)
	char := arr[i]
	if char >= 'A' && char <= 'Z' {
		arr[i] += 32
	} else if char >= 'a' && char <= 'z' {
		arr[i] -= 32
	} else {
		return
	}

	rec(arr, i+1, res)
}
