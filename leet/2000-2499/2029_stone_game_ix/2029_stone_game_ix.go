package _2029_stone_game_ix

// todo 2 read math solution with explanation.
// tptl passed medium (hard for me) #dfs #stone
func stoneGameIX(stones []int) bool {
	r := [3]int{}
	for _, s := range stones {
		r[s%3]++
	}

	min12 := min(r[1], r[2])
	if min12 != 0 {
		min12--
	}

	return dfs(r[0]%2, r[1]-min12, r[2]-min12, 3, true)
}

func dfs(r0, r1, r2, r int, isAlice bool) bool {
	if r == 0 || r0 < 0 || r1 < 0 || r2 < 0 {
		return true
	}

	if r0 == 0 && r1 == 0 && r2 == 0 {
		return !isAlice
	}

	return !dfs(r0-1, r1, r2, r, !isAlice) ||
		!dfs(r0, r1-1, r2, (r+1)%3, !isAlice) ||
		!dfs(r0, r1, r2-1, (r+2)%3, !isAlice)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
