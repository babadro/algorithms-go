package _1125_smallest_sufficient_team

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	var res []int

	reqSkills := make(map[string]bool)
	for _, skill := range req_skills {
		reqSkills[skill] = true
	}

	for {
		maxSkillSet := 0
		maxSkillSetIDx := -1

		for i, p := range people {
			skillSet := 0

			for _, skill := range p {
				if reqSkills[skill] {
					skillSet++
				}
			}

			if skillSet > maxSkillSet {
				maxSkillSet = skillSet
				maxSkillSetIDx = i
			}
		}

		for _, skill := range people[maxSkillSetIDx] {
			delete(reqSkills, skill)
		}

		res = append(res, maxSkillSetIDx)

		if len(reqSkills) == 0 {
			return res
		}
	}
}
