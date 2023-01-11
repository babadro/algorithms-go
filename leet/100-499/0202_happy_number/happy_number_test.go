package _202_happy_number

import (
	"fmt"
	"testing"
)

func Test_isHappy(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{19, true},
		{2, false},
		{1, true},
		{10, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := isHappy(tt.n); got != tt.want {
				t.Errorf("isHappy2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHappy(t *testing.T) {
	t.Log(isHappy(12))
}
