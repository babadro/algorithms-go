package _2141_maximum_running_time_of_N_computers

import "sort"

// https://leetcode.com/problems/maximum-running-time-of-n-computers/discuss/1692939/JavaC%2B%2BPython-Sort-Solution-with-Explanation-O(mlogm)
// tptl. passed. not mine. todo 2 difficult to understand
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
