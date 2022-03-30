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
	return recTopDown(0, make([]int, len(wealth)), wealth)
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

func findMaxStealBottomUp(wealth []int) int {
	if len(wealth) == 0 {
		return 0
	}

	dp := make([]int, len(wealth)+1) // +1 to handle the zero house
	dp[0], dp[1] = 0, wealth[0]

	// dp[] has one extra element to handle zero house
	for i := 1; i < len(wealth); i++ {
		dp[i+1] = utils.Max(wealth[i]+dp[i-1], dp[i])
	}

	return dp[len(wealth)]
}

func findMaxStealBottomUpOptimized(wealth []int) int {
	if len(wealth) == 0 {
		return 0
	}

	n1, n2 := 0, wealth[0]
	for i := 1; i < len(wealth); i++ {
		res := utils.Max(wealth[i]+n1, n2)
		n2, n1 = res, n2
	}

	return n2
}
