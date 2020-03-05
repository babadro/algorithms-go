package _1366_rank_teams_by_votes

import "testing"

func TestRankTeams(t *testing.T) {
	t.Log(rankTeams([]string{"ABC", "ACB", "ABC", "ACB", "ACB"}))
	t.Log(rankTeams([]string{"WXYZ", "XYZW"}))
	t.Log(rankTeams([]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"}))
	t.Log(rankTeams([]string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"}))
	t.Log(rankTeams([]string{"M", "M", "M", "M"}))
}
