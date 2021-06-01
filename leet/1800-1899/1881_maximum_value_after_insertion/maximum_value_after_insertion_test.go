package _1881_maximum_value_after_insertion

import (
	"fmt"
	"testing"
)

func Test_maxValue(t *testing.T) {
	tests := []struct {
		n    string
		x    int
		want string
	}{
		{"99", 9, "999"},
		{"-13", 2, "-123"},
		{"123", 1, "1231"},
		{"987", 4, "9874"},
		{"-987", 4, "-4987"},
		{"-487", 4, "-4487"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %d", tt.n, tt.x), func(t *testing.T) {
			if got := maxValue(tt.n, tt.x); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
