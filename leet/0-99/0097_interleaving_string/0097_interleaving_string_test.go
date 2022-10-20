package _097_interleaving_string

import "testing"

func Test_isInterleave(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
		{"abc", "abe", "abcabe", true},
		{"abc", "abe", "abeabc", true},
		{"a", "b", "a", false},
	}
	for _, tt := range tests {
		t.Run(tt.s1+"_"+tt.s2+"_"+tt.s3, func(t *testing.T) {
			if got := isInterleave(tt.s1, tt.s2, tt.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}
