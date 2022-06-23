package _1626_best_team_with_no_conflicts

import "sort"

type team struct {
	scores, ages []int
}

func (t team) Len() int {
	return len(t.scores)
}

func (t team) Less(i, j int) bool {
	if t.ages[i] != t.ages[j] {
		return t.ages[i] < t.ages[j]
	}

	return t.scores[i] < t.scores[j]
}

func (t team) Swap(i, j int) {
	t.ages[i], t.ages[j] = t.ages[j], t.ages[i]
	t.scores[i], t.scores[j] = t.scores[j], t.scores[i]
}

// tptl passed. todo 2 - bottom up solution should be much faster
func bestTeamScoreTopDown(scores []int, ages []int) int {
	t := team{
		scores: scores,
		ages:   ages,
	}
	sort.Sort(t)

	dp := make([][]int, len(scores))
	for i := range dp {
		dp[i] = make([]int, len(scores)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(dp, scores, 0, -1)
}

func recTopDown(dp [][]int, scores []int, curr, prev int) int {
	if curr == len(scores) {
		return 0
	}

	if dp[curr][prev+1] == -1 {
		res1 := 0
		if prev == -1 || scores[curr] >= scores[prev] {
			res1 = scores[curr] + recTopDown(dp, scores, curr+1, curr)
		}

		res2 := recTopDown(dp, scores, curr+1, prev)

		dp[curr][prev+1] = max(res1, res2)
	}

	return dp[curr][prev+1]
}

// correct, but tle
func bestTeamScoreBruteForce(scores []int, ages []int) int {
	t := team{
		scores: scores,
		ages:   ages,
	}
	sort.Sort(t)

	return recBruteForce(scores, ages, 0, -1)
}

func recBruteForce(scores, ages []int, curr, prev int) int {
	if curr == len(scores) {
		return 0
	}

	res1 := 0
	if prev == -1 || scores[curr] >= scores[prev] {
		res1 = scores[curr] + recBruteForce(scores, ages, curr+1, curr)
	}

	res2 := recBruteForce(scores, ages, curr+1, prev)

	return max(res1, res2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
