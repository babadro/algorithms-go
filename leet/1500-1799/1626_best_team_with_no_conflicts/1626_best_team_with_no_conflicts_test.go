package _1626_best_team_with_no_conflicts

import "testing"

func Test_bestTeamScoreBruteForce(t *testing.T) {
	tests := []struct {
		scores []int
		ages   []int
		want   int
	}{
		{[]int{1, 3, 5, 10, 15}, []int{1, 2, 3, 4, 5}, 34},
		{[]int{4, 5, 6, 5}, []int{2, 1, 2, 1}, 16},
		{[]int{1, 2, 3, 5}, []int{8, 9, 10, 1}, 6},
		{[]int{1, 3, 7, 3, 2, 4, 10, 7, 5}, []int{4, 5, 2, 1, 1, 2, 4, 1, 4}, 29},
		{tleScores1, tleAges1, 9864},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := bestTeamScoreTopDown(tt.scores, tt.ages); got != tt.want {
				t.Errorf("bestTeamScoreBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
