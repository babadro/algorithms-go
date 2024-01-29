package _2116_check_if_a_parentheses_string_can_be_valid

import "testing"

func Test_canBeValid(t *testing.T) {
	tests := []struct {
		s      string
		locked string
		want   bool
	}{
		//{
		//	"))()))", "010100", true,
		//},
		//{
		//	"()()", "0000", true,
		//},
		//{
		//	")", "0", false,
		//},
		//{
		//	"()", "11", true,
		//},
		{
			"())(()(()(())()())(())((())(()())((())))))(((((((())(()))))(",
			"100011110110011011010111100111011101111110000101001101001111",
			false,
		},
		{
			"((()(()()))()((()()))))()((()(()",
			"10111100100101001110100010001001",
			true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := canBeValid(tt.s, tt.locked); got != tt.want {
				t.Errorf("canBeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
