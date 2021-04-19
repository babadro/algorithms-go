package _069_Sqrt_x

import (
	"math"
	"sort"
)

func mySqrt(x int) int {
	res := sort.Search(x, func(i int) bool {
		return i*i > x
	})
	if res*res > x {
		res--
	}
	return res
}

func mySqrt2(x int) int {
	return int(math.Sqrt(float64(x)))
}
