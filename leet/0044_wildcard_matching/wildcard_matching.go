package _044_wildcard_matching

// https://leetcode.com/problems/wildcard-matching/discuss/347267/Go-8ms-DP-solution
// tptl. passed. best solution. hard. not mine.
func isMatch2(s string, p string) bool {
	mem := make([][]bool, len(s)+1)
	for i := range mem {
		mem[i] = make([]bool, len(p)+1)
	}

	mem[0][0] = true
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			mem[0][j] = mem[0][j-1]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				mem[i][j] = mem[i][j-1] || mem[i-1][j]
			} else if p[j-1] == '?' || p[j-1] == s[i-1] {
				mem[i][j] = mem[i-1][j-1]
			}
		}
	}

	return mem[len(s)][len(p)]
}

// bruteforce recursive. very slow. tle
func isMatchBruteForce(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	if len(p) == 0 {
		return false
	}

	char, nextP := p[0], p[1:]

	if char >= 'a' && char <= 'z' {
		if len(s) == 0 || s[0] != char {
			return false
		}

		return isMatch(s[1:], nextP)
	}

	if char == '?' {
		if len(s) == 0 {
			return false
		}

		return isMatch(s[1:], nextP)
	}

	// char == *
	if len(nextP) == 0 {
		return true
	}

	for i := 0; i <= len(s); i++ {
		if isMatch(s[i:], nextP) {
			return true
		}
	}

	return false
}

// tle
func isMatch(s string, p string) bool {
	mem := make(map[key]bool)

	return match(s, p, mem)
}

type key struct {
	s, p string
}

func match(s string, p string, mem map[key]bool) bool {
	k := key{s, p}

	if len(s) == 0 && len(p) == 0 {
		mem[k] = true
		return true
	}

	if len(p) == 0 {
		mem[k] = false
		return false
	}

	char, nextP := p[0], p[1:]

	if char >= 'a' && char <= 'z' {
		if len(s) == 0 || s[0] != char {
			mem[k] = false
			return false
		}

		if success, ok := mem[key{s[1:], nextP}]; ok {
			return success
		}

		return match(s[1:], nextP, mem)
	}

	if char == '?' {
		if len(s) == 0 {
			mem[k] = false
			return false
		}

		if success, ok := mem[key{s[1:], nextP}]; ok {
			return success
		}

		return match(s[1:], nextP, mem)
	}

	// char == *
	if len(nextP) == 0 {
		return true
	}

	for i := 0; i <= len(s); i++ {
		if success, ok := mem[key{s[i:], nextP}]; success && ok {
			return success
		}

		if match(s[i:], nextP, mem) {
			return true
		}
	}

	mem[k] = false
	return false
}
