package _242_valid_anagram

import (
	"testing"
)

func TestValidAnagram(t *testing.T) {
	cases := []struct {
		s1, s2   string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"", "", true},
		{"a", "a", true},
		{"car", "rat", false},
		{"arko", "arkoo", false},
		{"drrv", "rrdv", true},
		{"aacc", "ccac", false},
	}

	for i, c := range cases {
		if fact := isAnagram(c.s1, c.s2); fact != c.expected {
			t.Errorf("case#%d, s1 %s, s2 %s want isAnagram=%t, got %t", i+1, c.s1, c.s2, c.expected, fact)
		}
	}
}
