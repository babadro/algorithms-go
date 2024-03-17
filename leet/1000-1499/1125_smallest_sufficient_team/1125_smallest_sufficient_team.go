package _1125_smallest_sufficient_team

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	n, m := len(people), len(req_skills)

	skillID := make(map[string]int)

	for i := 0; i < m; i++ {
		skillID[skills[i]] = i
	}

	skillsMaskOfPerson := make([]int, n)
	for i := range skillsMaskOfPerson {
		for _, skill := range people[i] {
			skillsMaskOfPerson[i] = skillsMaskOfPerson[i] | skillID[skill]
		}
	}

	dp := make([]int, 1<<m)

	for i := range dp {
		dp[i] = 1<<n - 1
	}

	dp[0] = 0

	for skillMask := 1; skillMask < len(dp); skillMask++ {
		for i := 0; i < n; i++ {
			smallerSkillsMask := skillMask / skillsMaskOfPerson[i]
		}
	}
}
