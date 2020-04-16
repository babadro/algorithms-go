package _171_excel_sheet_column_number

import (
	"fmt"
	"math"
)

func titleToNumber(s string) int {
	p := float64(26)
	res := 0
	for i := 0; i < len(s); i++ {
		j := len(s) - i - 1
		res += int(math.Pow(p, float64(i))) * int(s[j]-65)
		fmt.Println(res)
	}
	return res
}
