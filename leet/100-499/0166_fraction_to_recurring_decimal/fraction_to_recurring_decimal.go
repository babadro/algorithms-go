package _166_fraction_to_recurring_decimal

import (
	"strconv"
)

// 100% 100% tptl. passed hard for me
func fractionToDecimal(numerator int, denominator int) string {
	b := make([]byte, 0, 18)
	if (numerator > 0 && denominator < 0) || (numerator < 0 && denominator > 0) {
		b = append(b, '-')
	}

	numerator, denominator = abs(numerator), abs(denominator)
	b = strconv.AppendInt(b, int64(numerator/denominator), 10)

	remainder := numerator % denominator
	if remainder == 0 {
		return string(b)
	}

	b = append(b, '.')

	m := make(map[int]int)
	idx, ok := 0, false
	for numerator = remainder; ; numerator %= denominator {
		if numerator == 0 {
			return string(b)
		}

		if idx, ok = m[numerator]; ok {
			break
		}

		m[numerator] = len(b)
		numerator = numerator * 10
		for numerator < denominator {
			b = append(b, '0')
			numerator *= 10
		}

		b = strconv.AppendInt(b, int64(numerator/denominator), 10)
	}

	periodLen := len(b) - idx
	period := make([]byte, periodLen)
	copy(period, b[idx:])
	b = b[:idx]
	b = append(b, '(')
	b = append(b, period...)
	b = append(b, ')')

	return string(b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

// doesn't work
func fractionToDecimal2(numerator int, denominator int) string {
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
