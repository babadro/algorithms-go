package _179_largest_number

import (
	"sort"
	"strconv"
)

// Doesn't work. Look testcase#10
func largestNumberFails(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	arr := make([]string, len(nums))
	for i, num := range nums {
		arr[i] = strconv.Itoa(num)
	}

	sort.Slice(arr, func(i, j int) bool {
		strI, strJ := arr[i], arr[j]
		if len(strI) == len(strJ) {
			return strI > strJ
		}

		k := 0
		for ; k < len(strI) && k < len(strJ); k++ {
			if strI[k] != strJ[k] {
				return strI[k] > strJ[k]
			}
		}

		if len(strI) < len(strJ) {
			return less(strI, strJ[k])
		}

		return !less(strJ, strI[k])
	})

	if arr[0][0] == '0' {
		return "0"
	}

	var res []byte
	for _, str := range arr {
		res = append(res, str...)
	}

	return string(res)
}

func less(str string, char byte) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != char {
			return str[i] > char
		}
	}
	return false
}
