package __coin_change

const divider = 1000007

func coinChangeBruteForce(denominations []int, total int) int {
	return recBruteForce(denominations, total, 0)
}

func recBruteForce(denominations []int, total, idx int) int {
	if total == 0 {
		return 1
	}

	if total < 0 || idx == len(denominations) {
		return 0
	}

	sum1 := recBruteForce(denominations, total-denominations[idx], idx)
	sum2 := recBruteForce(denominations, total, idx+1)

	return sum1 + sum2
}

// https://www.interviewbit.com/problems/coin-sum-infinite/
// passed. tptl
func coinChangeTopDown(denominations []int, total int) int {
	dp := make([][]int, len(denominations))
	for i := range dp {
		dp[i] = make([]int, total+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(dp, denominations, total, 0)
}

func recTopDown(dp [][]int, denominations []int, total, idx int) int {
	if total == 0 {
		return 1
	}

	if total < 0 || idx == len(denominations) {
		return 0
	}

	if dp[idx][total] == -1 {
		sum1 := recTopDown(dp, denominations, total-denominations[idx], idx) % divider
		sum2 := recTopDown(dp, denominations, total, idx+1) % divider

		dp[idx][total] = (sum1 + sum2) % divider
	}

	return dp[idx][total]
}

// passed. tptl
func coinChangeBottomUp(denominations []int, total int) int {
	dp := make([][]int, len(denominations))
	for i := range dp {
		dp[i] = make([]int, total+1)
	}

	for i := range dp {
		dp[i][0] = 1
	}

	for i := range denominations {
		for s := 1; s <= total; s++ {
			if i > 0 {
				dp[i][s] = dp[i-1][s]
			}

			if s >= denominations[i] {
				dp[i][s] += dp[i][s-denominations[i]]
			}

			dp[i][s] %= divider
		}
	}

	return dp[len(denominations)-1][total]
}
