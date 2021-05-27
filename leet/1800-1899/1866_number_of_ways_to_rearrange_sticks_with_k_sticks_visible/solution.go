package _866_number_of_ways_to_rearrange_sticks_with_k_sticks_visible

const mod = 1_000_000_007

var dp = [1001][1001]int{}

// todo 1 need to understand
// passed. tptl. best solution. hard
// https://leetcode.com/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible/discuss/1214875/Almost-AC-vs.-4-ms.
// explanation:
// https://leetcode.com/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible/discuss/1214679/C%2B%2B-Detailed-explanation-Clean-solution
func rearrangeSticksRecursive(n int, k int) int {
	if k == n {
		return 1
	}

	if k > 0 && n > k && dp[n][k] == 0 {
		dp[n][k] = ((n-1)*rearrangeSticksRecursive(n-1, k) + rearrangeSticksRecursive(n-1, k-1) + dp[n][k]) % mod
	}

	return dp[n][k]
}

// passed. simplier, but much slower
func rearrangeSticks(n int, k int) int {
	dp[0][0] = 1

	for i := 1; i <= 1000; i++ {
		for j := 1; j <= i; j++ {
			dp[i][j] = (dp[i-1][j-1] + (i-1)*dp[i-1][j]) % mod
		}
	}

	return dp[n][k]
}
