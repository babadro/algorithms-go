package _014_Longest_Common_Prefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		strs     []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{"1234"}, "1234"},

		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"dddd", "ddddd", "dddd"}, "dddd"},
		{[]string{"asdfasdf", "asd", "asdf"}, "asd"},
	}

	for i, c := range cases {
		if fact := longestCommonPrefix(c.strs); fact != c.expected {
			t.Errorf("case#%d, want %s, got %s", i+1, c.expected, fact)
		}
	}
}
