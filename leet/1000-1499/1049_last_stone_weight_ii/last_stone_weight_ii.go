package _1049_last_stone_weight_ii

// tptl. passed. best solution #dp medium
// todo 2 need to understand
func lastStoneWeightII(stones []int) int {
	totalSum := 0
	for _, stone := range stones {
		totalSum += stone
	}

	sum := totalSum / 2

	dp := make([]int, sum+1)
	for _, stone := range stones {
		for s := sum; s >= stone; s-- {
			dp[s] = max(dp[s], dp[s-stone]+stone)
		}
	}

	return totalSum - 2*dp[sum]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// passed. but too long.
func lastStoneWeightIIBottomUp(stones []int) int {
	totalSum := 0
	for _, stone := range stones {
		totalSum += stone
	}

	sum := totalSum / 2

	n := len(stones)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	for i := range dp {
		dp[i][0] = true
	}

	if stone := stones[0]; stone <= sum {
		dp[0][stone] = true
	}

	for i := 1; i < n; i++ {
		for s := 1; s <= sum; s++ {
			if dp[i-1][s] {
				dp[i][s] = true
			} else if s >= stones[i] {
				dp[i][s] = dp[i-1][s-stones[i]]
			}
		}
	}

	var sum1 int
	for s := sum; s >= 0; s-- {
		if dp[n-1][s] {
			sum1 = s
			break
		}
	}

	sum2 := totalSum - sum1

	return diff(sum1, sum2)
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
