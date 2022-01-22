package _2141_maximum_running_time_of_N_computers

import "sort"

func maxRunTime(n int, batteries []int) int64 {
	sort.Ints(batteries)
	sum := 0
	for _, b := range batteries {
		sum += b
	}

	k, lenBatteries := 0, len(batteries)
	for batteries[lenBatteries-1-k] > sum/(n-k) {
		sum -= batteries[lenBatteries-1-k]
		k++
	}

	return int64(sum / (n - k))
}
