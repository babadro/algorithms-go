package _1125_smallest_sufficient_team

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestSufficientTeam(t *testing.T) {
	tests := []struct {
		reqSkills []string
		people    [][]string
		want      []int
	}{
		//{
		//	[]string{"java", "nodejs", "reactjs"},
		//	[][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}},
		//	[]int{0, 2},
		//},
		//{
		//	[]string{"algorithms", "math", "java", "reactjs", "csharp", "aws"},
		//	[][]string{{"algorithms", "math", "java"}, {"algorithms", "math", "reactjs"}, {"java", "csharp", "aws"}, {"reactjs", "csharp"}, {"csharp", "math"}, {"aws", "java"}},
		//	[]int{1, 2},
		//},
		//{
		//	[]string{"1", "2", "3", "4", "5"},
		//	[][]string{{}, {}, {}, {}, {"2"}, {"2", "5"}, {"2", "3"}, {"1"}, {"4"}, {"1", "3"}},
		//	[]int{5, 8, 9},
		//},
		{
			skills,
			people,
			[]int{0, 15, 18, 29},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := smallestSufficientTeam2(tt.reqSkills, tt.people)

			sort.Ints(got)
			sort.Ints(tt.want)

			require.Equal(t, tt.want, got)
		})
	}
}
