package __palindromic_partitioning

import "github.com/babadro/algorithms-go/utils"

// see leetcode 132
func findMPPCuts(st string) int {
	return recBruteForce(st, 0, len(st)-1)
}

func recBruteForce(st string, startIDx, endIDx int) int {
	if startIDx >= endIDx || isPalindrome(st, startIDx, endIDx) {
		return 0
	}

	// at max, we need to cut the string into its "length-1" pieces
	minCuts := endIDx - startIDx
	for i := startIDx; i <= endIDx; i++ {
		if isPalindrome(st, startIDx, i) {
			minCuts = utils.Min(minCuts, 1+recBruteForce(st, i+1, endIDx))
		}
	}

	return minCuts
}

func findMPPCutsTopDown(st string) int {
	dp := make([][]int, len(st))
	for i := range dp {
		dp[i] = make([]int, len(st))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(st, dp, 0, len(st)-1)
}

func recTopDown(st string, dp [][]int, startIDx, endIDx int) int {
	if startIDx >= endIDx || isPalindrome(st, startIDx, endIDx) {
		return 0
	}

	if dp[startIDx][endIDx] == -1 {
		// at max, we need to cut the string into its "length-1" pieces
		minCuts := endIDx - startIDx
		for i := startIDx; i <= endIDx; i++ {
			if isPalindrome(st, startIDx, i) {
				minCuts = utils.Min(minCuts, 1+recBruteForce(st, i+1, endIDx))
			}
		}

		dp[startIDx][endIDx] = minCuts
	}

	return dp[startIDx][endIDx]
}

func isPalindrome(st string, start, end int) bool {
	for x, y := start, end; x < y; x, y = x+1, y-1 {
		if st[x] != st[y] {
			return false
		}
	}

	return true
}

func findMPPCutsBottomUp(st string) int {
	palindrome := make([][]bool, len(st))
	for i := 0; i < len(st); i++ {
		palindrome[i] = make([]bool, len(st))
		palindrome[i][i] = true
	}

	for startIDx := len(st) - 1; startIDx >= 0; startIDx-- {
		for endIDx := startIDx + 1; endIDx < len(st); endIDx++ {
			if st[startIDx] == st[endIDx] {
				if endIDx-startIDx == 1 || palindrome[startIDx+1][endIDx-1] {
					palindrome[startIDx][endIDx] = true
				}
			}
		}
	}

	cuts := make([]int, len(st))
	for startIDx := len(st) - 1; startIDx >= 0; startIDx-- {
		minCuts := len(st) // maximum cuts
		for endIDx := len(st) - 1; endIDx >= startIDx; endIDx-- {
			if palindrome[startIDx][endIDx] {
				if endIDx == len(st)-1 {
					minCuts = 0
				} else {
					minCuts = utils.Min(minCuts, 1+cuts[endIDx+1])
				}
			}
		}

		cuts[startIDx] = minCuts
	}

	return cuts[0]
}
