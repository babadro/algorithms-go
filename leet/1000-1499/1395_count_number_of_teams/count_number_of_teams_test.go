package _1395_count_number_of_teams

import (
	"fmt"
	"testing"
)

func Test_numTeams(t *testing.T) {
	tests := []struct {
		rating []int
		want   int
	}{
		{[]int{2, 5, 3, 4, 1}, 3},
		{[]int{2, 1, 3}, 0},
		{[]int{1, 2, 3, 4}, 4},
		{[]int{1, 2, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.rating), func(t *testing.T) {
			if got := numTeams2(tt.rating); got != tt.want {
				t.Errorf("numTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
