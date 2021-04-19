package _859_byddy_strings

import "testing"

func Test_buddyStrings(t *testing.T) {
	tests := []struct {
		name string
		A    string
		B    string
		want bool
	}{
		{"1", "ab", "ba", true},
		{"1", "ab", "ab", false},
		{"1", "aa", "aa", true},
		{"1", "aaaaaaabc", "aaaaaaacb", true},
		{"1", "", "", false},
		{"1", "a", "a", false},
		{"1", "bca", "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.A, tt.B); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
