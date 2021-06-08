package _1888_miminum_number_of_flips_to_make_the_binary_string_alternating

import "math"

// doesn't work
func minFlips2(s string) int {
	if len(s) == 1 {
		return 0
	}

	b := []byte(s)
	zeroFirstDist, oneFirstDist := dist(b)
	minDist := min(zeroFirstDist, oneFirstDist)
	if minDist == 0 {
		return 0
	}

	for i := 0; i < len(s); i++ {
		first := b[0]
		b = b[1:]
		b = append(b, first)

		if first == '1' {
			zeroFirstDist--
		} else {
			oneFirstDist--
		}

		zeroFirstDist, oneFirstDist = oneFirstDist, zeroFirstDist

		if len(s)%2 == 0 {
			if first == '1' {
				oneFirstDist++
			} else {
				zeroFirstDist++
			}
		} else {
			if first == '1' {
				zeroFirstDist++
			} else {
				oneFirstDist++
			}
		}

		if curr := min(zeroFirstDist, oneFirstDist); curr < minDist {
			minDist = curr
		}

		if minDist == 0 {
			return 0
		}
	}

	return minDist
}

func dist(b []byte) (int, int) {
	zeroFirst, oneFirst := 0, 0
	for i, char := range b {
		if i%2 == 0 {
			if char == '1' {
				oneFirst++
			} else {
				zeroFirst++
			}
		} else {
			if char == '1' {
				zeroFirst++
			} else {
				oneFirst++
			}
		}
	}

	return zeroFirst, oneFirst
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// tle, but correct
func minFlipsBruteForce(s string) int {
	if len(s) == 1 {
		return 0
	}

	b := []byte(s)
	minDist := math.MaxInt32
	for i := 0; i < len(s); i++ {
		b = append(b, b[0])
		b = b[1:]

		if curDist := distBruteForce(b); curDist < minDist {
			minDist = curDist
		}

		if minDist == 0 {
			return 0
		}
	}

	return minDist
}

func distBruteForce(b []byte) int {
	zeroFirst, oneFirst := 0, 0
	for i, char := range b {
		if i%2 == 0 {
			if char == '1' {
				oneFirst++
			} else {
				zeroFirst++
			}
		} else {
			if char == '1' {
				zeroFirst++
			} else {
				oneFirst++
			}
		}
	}

	if zeroFirst < oneFirst {
		return zeroFirst
	}

	return oneFirst
}
