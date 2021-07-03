package _134_gas_station

// tptl. passed. todo 2 need to understand
// https://leetcode.com/problems/gas-station/discuss/155280/Full-explanation
func canCompleteCircuit2(gas []int, cost []int) int {
	start, cur, account := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		cur += gas[i] - cost[i]
		if cur < 0 {
			start = i + 1
			account += cur
			cur = 0
		}
	}

	if cur+account >= 0 {
		return start
	}

	return -1
}
