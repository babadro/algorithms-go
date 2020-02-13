package _279_perfect_squares

import (
	"math"
)

func numSquares(n int) int {
	sqrt := math.Sqrt(float64(n))
	if sqrt == float64(int(sqrt)) {
		return 1
	}

	return 0
}

func perfectSquareNums(upperBound int) (seq []int) {
	num := 1
	for {
		perfectSquareNum := num * num
		if perfectSquareNum > upperBound {
			break
		}
		seq = append(seq, perfectSquareNum)
		num++
	}
	return seq
}
