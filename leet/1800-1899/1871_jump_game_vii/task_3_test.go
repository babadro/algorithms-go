package _1871_jump_game_vii

import (
	"testing"
)

func Test_canReach(t *testing.T) {
	tests := []struct {
		s       string
		minJump int
		maxJump int
		want    bool
	}{
		{"011010", 2, 3, true},
		{"01101110", 2, 3, false},
		{"011110", 1, 2, false},
		{"00", 1, 1, true},
		{tle1String, tle1MinJump, tle1MaxJump, true},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := canReach(tt.s, tt.minJump, tt.maxJump); got != tt.want {
				t.Errorf("canReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
