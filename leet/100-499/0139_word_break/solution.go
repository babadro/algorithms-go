package _139_word_break

// https://leetcode.com/problems/word-break/discuss/448646/Go-golang-0ms-solution
// passed. tptl medium #string
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}

	mem := make([]bool, len(s)+1)
	mem[0] = true
	for i := 1; i < len(mem); i++ {
		for j := i - 1; j >= 0; j-- {
			if mem[j] && dict[s[j:i]] {
				mem[i] = true

				break
			}
		}
	}

	return mem[len(mem)-1]
}
