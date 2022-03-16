package _2188_minimum_time_to_finish_the_race

import "math"

// https://leetcode.com/problems/minimum-time-to-finish-the-race/discuss/1802498/Pretreatment-%2B-DP
// passed #hard tptl #dp
func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	dp, best, maxLaps := make([]int, 1001), make([]int, 1001), 0

	for _, t := range tires {
		lapTime, time := t[0], t[0]
		for lap := 1; lap <= numLaps && lapTime < t[0]+changeTime; lap++ {
			maxLaps = max(maxLaps, lap)
			if best[lap] == 0 || best[lap] > time {
				best[lap] = time
			}

			lapTime *= t[1]
			time += lapTime
		}
	}

	return dfs(numLaps, changeTime, maxLaps, dp, best)
}

func dfs(laps, changeTime, maxLaps int, dp, best []int) int {
	if laps == 0 {
		return -changeTime
	}

	if dp[laps] == 0 {
		dp[laps] = math.MaxInt64
		for i := 1; i <= min(laps, maxLaps); i++ {
			dp[laps] = min(dp[laps], best[i]+changeTime+dfs(laps-i, changeTime, maxLaps, dp, best))
		}
	}

	return dp[laps]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
