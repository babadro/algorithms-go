package __house_thief

import "github.com/babadro/algorithms-go/utils"

func findMaxSteal(wealth []int) int {
	return recBruteForce(0, wealth)
}

func recBruteForce(idx int, wealth []int) int {
	if idx >= len(wealth) {
		return 0
	}

	res1 := wealth[idx] + recBruteForce(idx+2, wealth)
	res2 := recBruteForce(idx+1, wealth)

	return utils.Max(res1, res2)
}

func findMaxStealTopDown(wealth []int) int {
	return recTopDown(0, make([]int, len(wealth)+1), wealth)
}

func recTopDown(idx int, dp, wealth []int) int {
	if idx >= len(wealth) {
		return 0
	}

	if dp[idx] == 0 {
		res1 := wealth[idx] + recTopDown(idx+2, dp, wealth)
		res2 := recTopDown(idx+1, dp, wealth)

		dp[idx] = utils.Max(res1, res2)
	}

	return dp[idx]
}
