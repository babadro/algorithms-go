package _1323_Maximum_69_Number

import "strconv"

func maximum69Number(num int) int {
	str := []byte(strconv.Itoa(num))
	for i := range str {
		if str[i] == '6' {
			str[i] = '9'
			break
		}
	}
	res, _ := strconv.Atoi(string(str))
	return res
}
