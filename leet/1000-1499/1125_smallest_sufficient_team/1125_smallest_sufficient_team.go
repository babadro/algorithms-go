package _1125_smallest_sufficient_team

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	var bestRes []int

	reqSkills := make(map[string]bool)
	for _, skill := range req_skills {
		reqSkills[skill] = true
	}

	rec(reqSkills, -1, nil, &bestRes, people)

	return bestRes
}

func rec(reqSkills map[string]bool, idx int, curr []int, bestRes *[]int, people [][]string) {
	if idx != -1 {
		if len(*bestRes) > 0 && len(curr)+1 == len(*bestRes) {
			return
		}

		for _, skill := range people[idx] {
			delete(reqSkills, skill)
		}

		curr = append(curr, idx)
	}

	if len(reqSkills) == 0 {
		*bestRes = make([]int, len(curr))
		copy(*bestRes, curr)

		if idx != -1 {
			for _, skill := range people[idx] {
				reqSkills[skill] = true
			}
		}

		return
	}

	var bestPeople []int
	maxSkillset := 0

	for i, p := range people {
		skillCount := 0

		for _, skill := range p {
			if reqSkills[skill] {
				skillCount++
			}
		}

		if skillCount == 0 {
			continue
		}

		if skillCount == maxSkillset {
			bestPeople = append(bestPeople, i)
		} else if skillCount > maxSkillset {
			bestPeople = []int{i}
			maxSkillset = skillCount
		}
	}

	for _, peopleIDx := range bestPeople {
		rec(reqSkills, peopleIDx, curr, bestRes, people)
	}

	if idx != -1 {
		for _, skill := range people[idx] {
			reqSkills[skill] = true
		}
	}
}
