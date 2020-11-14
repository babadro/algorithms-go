package _290_word_pattern

import "testing"

func Test_wordPattern(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		s       string
		want    bool
	}{
		{"1", "abba", "dog cat cat dog", true},
		{"1", "abba", "dog cat cat fish", false},
		{"1", "aaaa", "dog cat cat dog", false},
		{"1", "abba", "dog dog dog dog", false},
		{"1", "abba", "dog word word dog", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern2(tt.pattern, tt.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
