package _322_coin_change

func coinChangeBruteForce2(coins []int, amount int) int {
	return recBruteForce(coins, 0, 0, amount)
}

func recBruteForce(coins []int, cur, idx, amount int) int {
	if amount == 0 {
		return cur
	}

	if idx == len(coins) {
		return 0
	}

	sum1 := 0
	if amount >= coins[idx] {
		sum1 = recBruteForce(coins, idx, cur+1, amount-coins[idx])
	}

	sum2 := recBruteForce(coins, idx+1, cur, amount)

	return max(sum1, sum2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
