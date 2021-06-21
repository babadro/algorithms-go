package _1904_the_number_of_full_rounds_you_have_played

import (
	"fmt"
	"testing"
)

func Test_numberOfRounds(t *testing.T) {
	tests := []struct {
		startTime  string
		finishTime string
		want       int
	}{
		{"12:01", "12:44", 1},
		{"20:00", "06:00", 40},
		{"00:00", "23:59", 95},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s", tt.startTime, tt.finishTime), func(t *testing.T) {
			if got := numberOfRounds2(tt.startTime, tt.finishTime); got != tt.want {
				t.Errorf("numberOfRounds() = %v, want %v", got, tt.want)
			}
		})
	}
}
