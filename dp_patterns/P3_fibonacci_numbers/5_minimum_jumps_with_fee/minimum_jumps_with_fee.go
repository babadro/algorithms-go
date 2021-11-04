package __minimum_jumps_with_fee

import (
	"github.com/babadro/algorithms-go/utils"
)

func findMinFeeBruteForce(fee []int) int {
	return recBruteForce(fee, 0)
}

func recBruteForce(fee []int, idx int) int {
	if idx >= len(fee) {
		return 0
	}

	res1, res2, res3 := recBruteForce(fee, idx+1), recBruteForce(fee, idx+2), recBruteForce(fee, idx+3)

	return fee[idx] + utils.Min3(res1, res2, res3)
}

func findMinFeeTopDown(fee []int) int {
	dp := make([]int, len(fee))

	return recTopDown(dp, fee, 0)
}

func recTopDown(dp []int, fee []int, idx int) int {
	if idx >= len(fee) {
		return 0
	}

	if dp[idx] == 0 {
		res1, res2, res3 := recBruteForce(fee, idx+1), recBruteForce(fee, idx+2), recBruteForce(fee, idx+3)

		dp[idx] = fee[idx] + utils.Min3(res1, res2, res3)
	}

	return dp[idx]
}

func findMinFeeBottomUp(fee []int) int {
	dp := make([]int, len(fee)+1)

	dp[0] = 0
	dp[1], dp[2], dp[3] = fee[0], fee[0], fee[0]

	for i := 3; i < len(fee); i++ {
		dp[i+1] = utils.Min3(fee[i]+dp[i], fee[i-1]+dp[i-1], fee[i-2]+dp[i-2])
	}

	return dp[len(fee)]
}

func findMinFeeOptimized(fee []int) int {
	dp := [4]int{0, fee[0], fee[0], fee[0]}

	for i := 3; i < len(fee); i++ {
		next := utils.Min3(fee[3]+dp[3], fee[2]+dp[2], fee[1]+dp[1])
		dp[0], dp[1], dp[2] = dp[1], dp[2], dp[3]
		dp[3] = next
	}

	return dp[3]
}
