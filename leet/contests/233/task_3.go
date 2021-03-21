package _33

func maxValue(n int, index int, maxSum int) int {
	l, r := 1, maxSum
	for l < r {
		mid := (l + r) / 2

		if curr := sum(n, index, mid); curr > maxSum {
			r = mid - 1
		} else if curr == maxSum {
			return mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func sum(n, index, num int) int {
	sumLeft, sumRight := 0, 0

	startLeft := num - index
	if startLeft < 1 {
		sumLeft = sumArProg(1, num, num) + index - num + 1
	} else {
		sumLeft = sumArProg(num-index, num, index+1)
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
