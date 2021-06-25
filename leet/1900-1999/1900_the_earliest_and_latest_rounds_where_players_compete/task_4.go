package _1900_the_earliest_and_latest_rounds_where_players_compete

import "math"

// todo 1 more effective solution with memorization
// https://leetcode.com/problems/the-earliest-and-latest-rounds-where-players-compete/discuss/1268586/Java-DP-with-memory
func earliestAndLatest2(n int, firstPlayer int, secondPlayer int) []int {
	return nil
}

// https://leetcode.com/problems/the-earliest-and-latest-rounds-where-players-compete/discuss/1268539/Recursion-(500ms)-%2B-Memo-(0ms)
var minR int
var maxR int

// todo 1 need to understand
func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	minR, maxR = math.MaxInt32, math.MinInt32
	dfs((1<<n)-1, 1, 0, 27, firstPlayer-1, secondPlayer-1)

	return []int{minR, maxR}
}

func dfs(mask, round, i, j, first, second int) {
	if i >= j {
		dfs(mask, round+1, 0, 27, first, second)
	} else if mask&(1<<i) == 0 {
		dfs(mask, round, i+1, j, first, second)
	} else if mask&(1<<j) == 0 {
		dfs(mask, round, i, j-1, first, second)
	} else if i == first && j == second {
		minR = min(minR, round)
		maxR = max(maxR, round)
	} else {
		if i != first && i != second {
			dfs(mask^(1<<i), round, i+1, j-1, first, second)
		}
		if j != first && j != second {
			dfs(mask^(1<<j), round, i+1, j-1, first, second)
		}
	}
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
