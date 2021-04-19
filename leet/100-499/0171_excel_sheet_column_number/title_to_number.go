package _171_excel_sheet_column_number

import (
	"math"
)

func titleToNumber(s string) int {
	var res int
	for _, char := range s {
		res = res*26 + int(char-'A'+1)
	}
	return res
}

func titleToNumber2(s string) int {
	p := float64(26)
	res := 0
	for i := 0; i < len(s); i++ {
		j := len(s) - i - 1
		res += int(math.Pow(p, float64(i))) * int(s[j]-64)
	}
	return res
}
