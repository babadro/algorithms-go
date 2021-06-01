package _1883_minimum_skips_to_arrive_at_meeting_on_time

// https://leetcode.com/problems/minimum-skips-to-arrive-at-meeting-on-time/discuss/1239838/JavaC%2B%2BPython-DP-Solution
// todo 1 need to understand
func minSkips(dist []int, speed int, hoursBefore int) int {
	n := len(dist)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		for j := n; j >= 0; j-- {
			dp[j] += dist[i]

			if i < n-1 {
				dp[j] = (dp[j] + speed - 1) / speed * speed
			}

			if j > 0 {
				dp[j] = min(dp[j], dp[j-1]+dist[i])
			}
		}
	}

	for k := 0; k < n; k++ {
		if dp[k] <= speed*hoursBefore {
			return k
		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
