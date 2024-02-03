package _2131_longest_palindrome_by_concatenating_two_letter_words

import "testing"

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{words: []string{"lc", "cl", "gg"}, want: 6},
		{words: []string{"ab", "ty", "yt", "lc", "cl", "ab"}, want: 8},
		{words: []string{"cc", "ll", "xx"}, want: 2},
		{words: []string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}, want: 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.words); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
