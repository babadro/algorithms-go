package _1601_maximum_number_of_achievable_transfer_requests

// #bnsrg hard passed
// todo 2 check bitmask solution
func maximumRequests(n int, requests [][]int) int {
	buildings := make([]int, n)

	var res int

	rec(buildings, requests, 0, 0, &res)

	return res
}

func rec(buildings []int, requests [][]int, idx, count int, res *int) {
	if idx == len(requests) {
		for _, balance := range buildings {
			if balance != 0 {
				return
			}
		}

		*res = max(*res, count)

		return
	}

	rec(buildings, requests, idx+1, count, res)

	from, to := requests[idx][0], requests[idx][1]

	buildings[from]--
	buildings[to]++

	rec(buildings, requests, idx+1, count+1, res)

	buildings[from]++
	buildings[to]--
}
