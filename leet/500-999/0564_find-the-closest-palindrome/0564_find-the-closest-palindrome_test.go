package _564_find_the_closest_palindrome

import "testing"

func Test_nearestPalindromic(t *testing.T) {
	tests := []struct {
		n    string
		want string
	}{
		{"123", "121"},
		{"1", "0"},
		{"1234", "1221"},
		{"999", "1001"},
		{"1000", "999"},
		{"12932", "12921"},
		{"99800", "99799"},
		{"12120", "12121"},
		{"10", "9"},
		{"11", "9"},
		{"12", "11"},
		{"19", "22"},
	}
	for _, tt := range tests {
		t.Run(tt.n, func(t *testing.T) {
			if got := nearestPalindromic(tt.n); got != tt.want {
				t.Errorf("nearestPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}
