package _1544_make_the_string_great

import "testing"

func Test_makeGood(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"leEeetcode", "leetcode"},
		{"abBAcC", ""},
		{"s", "s"},
		{"aa", "aa"},
		{"Aa", ""},
		{"abcd", "abcd"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := makeGood(tt.s); got != tt.want {
				t.Errorf("makeGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
