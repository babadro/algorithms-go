package _1366_rank_teams_by_votes

import "sort"

func rankTeams(votes []string) string {
	length := len(votes[0])
	sortedTeams := []int32(votes[0])
	places := make(map[int32][]int)

	for _, vote := range votes {
		for j, team := range vote {
			if _, ok := places[team]; !ok {
				places[team] = make([]int, length)
			}
			places[team][j]++
		}
	}

	sort.Slice(sortedTeams, func(i, j int) bool {
		for k := 0; k < length; k++ {
			countI, countJ := places[sortedTeams[i]][k], places[sortedTeams[j]][k]
			if countI == countJ {
				continue
			}
			return countI > countJ
		}
		return sortedTeams[i] < sortedTeams[j]
	})

	return string(sortedTeams)
}
