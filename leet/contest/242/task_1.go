package _42

import "github.com/babadro/algorithms-go/utils"

// passed. easy. todo 3 find shorter solution
func checkZeroOnes(s string) bool {
	s += " "
	counterOnes := 0
	curr := 0
	for _, char := range s {
		if char != '1' {
			counterOnes = utils.Max(curr, counterOnes)
			curr = 0

			continue
		} else if char == '1' {
			curr++
		}
	}

	curr = 0
	counterZeroes := 0
	for _, char := range s {
		if char != '0' {
			counterZeroes = utils.Max(curr, counterZeroes)
			curr = 0

			continue
		} else if char == '0' {
			curr++
		}
	}

	return counterOnes > counterZeroes
}
