package _33

func maxValue(n int, index int, maxSum int) int {
	l, r := 1, 1_000_000_000
	for l < r {
		mid := (l + r) / 2

		if sum(n, index, mid) > maxSum {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func sum(n, index, num int) int {
	return sumArProg(num-index, num, index+1) +
		sumArProg(num, num-(n-1-index), n-index) -
		num
}

func sumArProg(a1, an, n int) int {
	return n * (a1 + an) / 2
}
