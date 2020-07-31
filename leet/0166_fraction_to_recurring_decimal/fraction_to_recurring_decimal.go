package _166_fraction_to_recurring_decimal

import (
	"strconv"
)

// TODO 1
func fractionToDecimal(numerator int, denominator int) string {
	b := make([]byte, 0, 18)
	leftPart := numerator / denominator
	b = strconv.AppendInt(b, int64(leftPart), 10)

	remainder := numerator % denominator
	if remainder == 0 {
		return string(b)
	}

	b = append(b, '.')
	rightPart := float64(remainder) / float64(denominator)

	rightPartStr := strconv.FormatFloat(rightPart, 'f', -1, 64)[2:]
	maxPeriodLen := len(rightPartStr) / 2

Loop:
	for periodLen := 1; periodLen <= maxPeriodLen; periodLen++ {
		periodsCount := len(rightPartStr) / periodLen
		period := rightPartStr[:periodLen]
		for j := 1; j < periodsCount; j++ {
			from := j * periodLen
			to := from + periodLen
			if rightPartStr[from:to] != period {
				continue Loop
			}
			if j == periodsCount-1 {
				b = append(b, '(')
				b = append(b, period...)
				b = append(b, ')')

				return string(b)
			}
		}
	}

	b = append(b, rightPartStr...)
	return string(b)
}
