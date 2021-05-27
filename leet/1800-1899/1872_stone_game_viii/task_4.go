package _1872_stone_game_viii

// todo 2 need to understand passed. tptl. best solution. hard
// https://leetcode.com/problems/stone-game-viii/discuss/1224872/Top-Down-and-Bottom-Up
func stoneGameVIII(stones []int) int {
	n := len(stones)
	for i := 1; i < n; i++ {
		stones[i] += stones[i-1]
	}

	res := stones[n-1]
	for i := n - 2; i > 0; i-- {
		res = max(res, stones[i]-res)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
