package _125_valid_palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"1", "A man, a plan, a canal: Panama", true},
		{"1", "race a car", false},
		{"1", "", true},
		{"1", "a", true},
		{"1", "aA", true},
		{"1", "aBba", true},
		{"1", "a...................Bba", true},
		{"1", "......,,,,,,,1a", false},
		{"1", "12321a", false},
		{"1", "12321!!!!!!", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
