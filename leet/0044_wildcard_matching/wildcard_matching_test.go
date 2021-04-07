package _044_wildcard_matching

import "testing"

func Test_isMatch(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "*", true},
		{"cb", "?a", false},
		{"abceb", "*a*b", true},
		{"acdcb", "a*c?b", false},
		{"", "*", true},
		{"a", "b", false},
		{"a", "?", true},
		{"", "?", false},
		{"abcd", "****", true},
		{"adddasdfasdfasdfasdfasdfadsfasdfasdfhah", "a*?*?*?*?*?*?*?*?*?*?*?*?h", true},
		{"aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba", "a*******b", false},
	}
	for _, tt := range tests {
		t.Run(tt.s+" "+tt.p, func(t *testing.T) {
			if got := isMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

// "aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba"
//"a*******b"
