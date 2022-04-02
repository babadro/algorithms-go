package __palindromic_partitioning

import "github.com/babadro/algorithms-go/utils"

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

func isPalindrome(st string, start, end int) bool {
	for x, y := start, end; x < y; x, y = x+1, y-1 {
		if st[x] != st[y] {
			return false
		}
	}

	return true
}
