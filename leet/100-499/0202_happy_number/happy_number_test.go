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
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := isHappy(tt.n); got != tt.want {
				t.Errorf("isHappy2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcSquareSum(t *testing.T) {
	t.Log(replace(123467))
}

func TestIsHappy(t *testing.T) {
	t.Log(isHappy2(12))
}
