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

func dynamicCoinChanging(coins []int, V int) int {
	m := len(coins)
	table := make([]int, V+1)
	table[0] = 0

	for i := 1; i <= V; i++ {
		table[i] = math.MaxInt32
	}

	for i := 1; i <= V; i++ {
		for j := 0; j < m; j++ {
			if coins[j] <= i {
				subRes := table[i-coins[j]]
				if subRes != math.MaxInt32 && subRes+1 < table[i] {
					table[i] = subRes + 1
				}
			}
		}
	}
	return table[V]
}
