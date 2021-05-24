package _1871_jump_game_vii

// passed, but it's slow solution
func canReachSlow(s string, minJump int, maxJump int) bool {
	n := len(s)
	last := n - 1
	if s[last] == '1' {
		return false
	}

	memo := make([]bool, n)
	memo[last] = true
	for i := last - 1; i >= 0; i-- {
		char := s[i]
		if char == '1' {
			continue
		}

		for j := i + minJump; j < n && j <= i+maxJump; j++ {
			if memo[j] {
				memo[i] = true
				break
			}
		}
	}

	return memo[0]
}
