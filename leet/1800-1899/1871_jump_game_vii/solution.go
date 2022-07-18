package _1871_jump_game_vii

import "github.com/babadro/algorithms-go/utils"

// tptl. passed. best solution. medium (this solution hard for me)
// todo 2 look at sliding window solution
func canReach(s string, minJump int, maxJump int) bool {
	b := []byte(s)
	n := len(b)
	for i, j := 0, 0; i < n; i++ {
		if i == 0 || b[i] == '2' {
			for j = utils.Max(j, i+minJump); j <= utils.Min(n-1, i+maxJump); j++ {
				if b[j] == '0' {
					b[j] = '2'
				}
			}
		}
	}

	return b[n-1] == '2'
}
