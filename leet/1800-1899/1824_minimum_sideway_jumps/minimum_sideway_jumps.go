package _1824_minimum_sideway_jumps

import (
	"github.com/babadro/algorithms-go/utils"
	"math"
)

// not mine. passed. medium. tptl. todo 1 need to outline
// c++ from here https://leetcode.com/problems/minimum-sideway-jumps/discuss/1152455/JavaC%2B%2BPython-DP-%2B-Greedy-Bonus
func minSideJumpsDP(obstacles []int) int {
	dp := [4]int{0, 1, 0, 1}
	for _, obs := range obstacles {
		dp[obs] = math.MaxInt32
		for j := 1; j <= 3; j++ {
			if j != obs {
				jumps1, jumps2, jumps3 := dp[1], dp[2], dp[3]
				if j != 1 {
					jumps1 += 1
				}
				if j != 2 {
					jumps2 += 1
				}
				if j != 3 {
					jumps3 += 1
				}

				dp[j] = utils.Min3(jumps1, jumps2, jumps3)
			}
		}
	}

	return utils.Min3(dp[1], dp[2], dp[3])
}

// todo 2 doesn't work. Can't understand the solution
// https://leetcode.com/problems/minimum-sideway-jumps/discuss/1152455/JavaC%2B%2BPython-DP-%2B-Greedy-Bonus
func minSideJumpsGreedy(obstacles []int) int {
	res, cur := 0, 2
	for i := 0; i < len(obstacles)-1; i++ {
		if cur == obstacles[i+1] {
			res++
			for b := [4]bool{}; i < len(obstacles); i++ {
				b[obstacles[i]] = true
				if b[1] && b[2] && b[3] {
					break
				}
			}

			cur = obstacles[i]
			i -= 2
		}
	}

	return res
}
