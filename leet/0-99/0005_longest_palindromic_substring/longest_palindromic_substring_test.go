package _005_longest_palindromic_substring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"bab", "bab"},
		{"baba", "bab"},
		//	{"", ""},
		{"1", "1"},
		{"asde", "a"},
		{"asded", "ded"},
	}

	for i, c := range cases {
		if fact := longestPalindrome3(c.input); fact != c.expected {
			t.Errorf("case#%d, want %s, got %s", i+1, c.expected, fact)
		}
	}
}

func TestLongestPalindrome2(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"bab", "bab"},
		{"baba", "aba"},
		{"", ""},
		{"1", "1"},
		{"asde", "e"},
		{"asded", "ded"},
		{"babad", "bab"},
	}

	for _, c := range cases {
		require.Equal(t, c.expected, longestPalindrome2(c.input))
	}
}

/*
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
*/
