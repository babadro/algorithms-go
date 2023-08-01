package _975_odd_even_jump

import "math"

const (
	goodOdd  = byte(1) << 0
	goodEven = byte(1) << 1
)

func oddEvenJumps(arr []int) int {
	counter := 1

	indices := make([]byte, len(arr))
	indices[len(indices)-1] = goodEven | goodOdd

	for i := len(arr) - 2; i >= 0; i-- {
		idx := oddJump(i, arr)
		if idx != -1 && indices[idx]&goodEven != 0 {
			indices[i] = goodOdd
			counter++
		}

		idx = evenJump(i, arr)
		if idx != -1 && indices[idx]&goodOdd != 0 {
			indices[i] |= goodEven
		}
	}

	return counter
}

func oddJump(i int, arr []int) int {
	res := -1
	jumpFrom := arr[i]
	diff := math.MaxInt64 - jumpFrom
	for i = i + 1; i < len(arr); i++ {
		num := arr[i]
		if num < jumpFrom {
			continue
		}

		if d := num - jumpFrom; d < diff {
			res = i
			diff = d
		}

		if diff == 0 {
			break
		}
	}

	return res
}

func evenJump(i int, arr []int) int {
	res := -1
	jumpFrom := arr[i]
	diff := -1
	for i = i + 1; i < len(arr); i++ {
		num := arr[i]
		if num > jumpFrom {
			continue
		}

		if d := jumpFrom - num; d > diff {
			res = i
			diff = d
		}
	}

	return res
}
