package _003_lenOfLongestSubStr

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"dvdf", 3},
		{"abcdecagifn", 8},
	}

	for i, c := range cases {
		if fact := lengthOfLongestSubstring4(c.input); fact != c.expected {
			t.Errorf("case#%d error. want %d, got %d", i+1, c.expected, fact)
		}
	}
}
