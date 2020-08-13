package _134_gas_station

// 20% 100%. This is a brutforce
// TODO 2 find quickier solution in discuss
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
Loop:
	for start := 0; start < n; start++ {
		tank := 0
		for j := 0; j < n; j++ {
			idx := (start + j) % n
			tank += gas[idx]
			tank -= cost[idx]
			if tank < 0 {
				continue Loop
			}
		}
		return start
	}
	return -1
}
