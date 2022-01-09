package _2127

import "testing"

func Test_maximumInvitations(t *testing.T) {
	tests := []struct {
		favorite []int
		want     int
	}{
		{[]int{2, 2, 1, 2}, 3},
		{[]int{1, 2, 0}, 3},
		{[]int{3, 0, 1, 4, 1}, 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maximumInvitations(tt.favorite); got != tt.want {
				t.Errorf("maximumInvitations() = %v, want %v", got, tt.want)
			}
		})
	}
}
