package __unique_generalized_abbreviations

import "strconv"

// leetcode (premium) https://leetcode.com/problems/generalized-abbreviation/
func generateGeneralizedAbbreviation(s string) []string {
	var res []string
	rec(0, 0, s, make([]byte, 0, len(s)), &res)

	return res
}

func rec(i, prev int, s string, cur []byte, res *[]string) {
	if i == len(s) {
		if prev > 0 {
			cur = append(cur, strconv.Itoa(prev)...)
		}

		*res = append(*res, string(cur))

		return
	}

	rec(i+1, prev+1, s, cur, res)

	if prev > 0 {
		cur = append(cur, strconv.Itoa(prev)...)
	}

	cur = append(cur, s[i])
	rec(i+1, 0, s, cur, res)
}
