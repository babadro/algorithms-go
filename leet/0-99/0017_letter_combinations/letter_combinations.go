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

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	d := make([]digit, 0, len(digits))
	for _, b := range digits {
		d = append(d, digit{letters: dict[b-'0']})
	}
	res := make([]string, 0, int(math.Pow(3, float64(len(digits)))))
	combination := make([]byte, 0, len(digits))
	for {
		combination = combination[:0]
		for i := 0; i < len(d); i++ {
			combination = append(combination, d[i].letters[d[i].idx])
		}
		res = append(res, string(combination))
		if !next(d) {
			break
		}
	}
	return res
}

func next(digits []digit) bool {
	for i := 0; i < len(digits); i++ {
		if digits[i].idx < len(digits[i].letters)-1 {
			digits[i].idx++
			return true
		}
		digits[i].idx = 0
	}
	return false
}
