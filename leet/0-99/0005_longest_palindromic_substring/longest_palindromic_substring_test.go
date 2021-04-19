package _005_longest_palindromic_substring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"bab", "bab"},
		{"baba", "bab"},
		{"", ""},
		{"1", "1"},
		{"asde", "a"},
		{"asded", "ded"},
	}

	for i, c := range cases {
		if fact := longestPalindrome(c.input); fact != c.expected {
			t.Errorf("case#%d, want %s, got %s", i+1, c.expected, fact)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"bab", true},
		{"baba", false},
		{"", true},
		{"1", true},
	}

	for i, c := range cases {
		if fact := isPalindrome([]rune(c.input)); fact != c.expected {
			t.Errorf("case#%d, want %t, got %t", i+1, c.expected, fact)
		}
	}
}
