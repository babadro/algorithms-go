package _017_letter_combinations

import (
	"math"
)

var dict = [10][]byte{
	2: {'a', 'b', 'c'},
	3: {'d', 'e', 'f'},
	4: {'g', 'h', 'i'},
	5: {'j', 'k', 'l'},
	6: {'m', 'n', 'o'},
	7: {'p', 'q', 'r', 's'},
	8: {'t', 'u', 'v'},
	9: {'w', 'x', 'y', 'z'},
}

type digit struct {
	idx     int
	letters []byte
}

// TODO 1 (it's like ammo generator)
func letterCombinations(digits string) []string {
	d := make([]digit, 0, len(digits))
	for _, b := range digits {
		d = append(d, digit{letters: dict[b-'0']})
	}
	res := make([]string, 0, int(math.Pow(3, float64(len(digits)))))
	for {
		str, ok := next(d)
		if !ok {
			break
		}
		res = append(res, str)
	}
	return res
}

func next(digits []digit) (string, bool) {
	for i := 0; i < len(digits); i++ {
		if digits[i].idx < len(digits[i].letters) {
			var res []byte
			for k := 0; k < len(digits); k++ {
				res = append(res, digits[k].letters[digits[k].idx])
			}
			//fmt.Println(string(res))
			digits[i].idx++
			return string(res), true
		} else {
			digits[i].idx = 0
		}
	}
	return "", false
}
