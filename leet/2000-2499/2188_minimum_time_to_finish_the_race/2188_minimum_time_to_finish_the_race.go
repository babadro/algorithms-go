package _2188_minimum_time_to_finish_the_race

import "math"

var dp, best [1001]int

var maxLaps = 0

// todo 1 finish it
func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
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

	return dfs(numLaps, changeTime)
}

func dfs(laps, changeTime int) int {
	if laps == 0 {
		return -changeTime
	}

	if dp[laps] != 0 {
		dp[laps] = math.MaxInt64
		for i := 1; i < min(laps, maxLaps); i++ {
			dp[laps] = min(dp[laps], best[i]+changeTime+dfs(laps-i, changeTime))
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
