package _1802_maximum_value_at_a_given_index_in_a_bounded_array

import (
	"github.com/babadro/algorithms-go/utils"
	"sort"
)

// passed. tptl. best solution. medium
// todo 1 need to understand
func maxValue2(n int, index int, maxSum int) int {
	f := func(num int) bool {
		leftLen := utils.Min(num-1, index+1)
		rightLen := utils.Min(num-1, n-index)

		leftSum := (2*num - leftLen - 1) * leftLen / 2
		rightSum := (2*num - rightLen - 1) * rightLen / 2

		return leftSum+rightSum-num+1+n > maxSum
	}

	return sort.Search(maxSum+1, f) - 1
}

// passed. mine
func maxValue(n int, index int, maxSum int) int {
	f := func(num int) bool {
		return sum(n, index, num) > maxSum
	}

	res := sort.Search(maxSum+1, f)

	return res - 1
}

func sum(n, index, num int) int {
	sumLeft, sumRight := 0, 0

	startLeft := num - index
	if startLeft < 1 {
		sumLeft = sumArProg(1, num, num) + index - num + 1
	} else {
		sumLeft = sumArProg(startLeft, num, index+1)
	}

	finishRight := num - (n - 1 - index)
	if finishRight < 1 {
		sumRight = sumArProg(num, 1, num) + (n - index - num)
	} else {
		sumRight = sumArProg(num, finishRight, n-index)
	}

	return sumLeft + sumRight - num
}

func sumArProg(a1, an, n int) int {
	return n * (a1 + an) / 2
}
