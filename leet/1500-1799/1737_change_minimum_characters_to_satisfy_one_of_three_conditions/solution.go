package _1737_change_minimum_characters_to_satisfy_one_of_three_conditions

// best solution
// todo 2 understand
// tptl
func minCharacters2(a string, b string) int {
	c1, c2 := make([]int, 26), make([]int, 26)
	for _, v := range a {
		c1[v-'a'] += 1
	}
	for _, v := range b {
		c2[v-'a'] += 1
	}
	m, n := len(a), len(b)
	ans := m + n
	for i := 0; i < 26; i++ {
		ans = min(ans, m+n-c1[i]-c2[i])
		if i > 0 {
			c1[i] += c1[i-1]
			c2[i] += c2[i-1]
		}
		if i < 25 {
			ans = min(ans, m-c1[i]+c2[i])
			ans = min(ans, n-c2[i]+c1[i])
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
