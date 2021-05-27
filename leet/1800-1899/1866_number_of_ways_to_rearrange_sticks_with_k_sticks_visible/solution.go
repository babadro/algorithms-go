package _866_number_of_ways_to_rearrange_sticks_with_k_sticks_visible

const mod = 1_000_000_007

var dp = [1001][1001]int{}

func rearrangeSticks(n int, k int) int {
	if k == n {
		return 1
	}

	if k > 0 && n > k && dp[n][k] == 0 {
		dp[n][k] = ((n-1)*rearrangeSticks(n-1, k) + rearrangeSticks(n-1, k-1) + dp[n][k]) % mod
	}

	return dp[n][k]
}
