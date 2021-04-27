package _839_longest_substring_of_al_vowels_in_order

import "testing"

func Test_longestBeautifulSubstring(t *testing.T) {
	tests := []struct {
		word string
		want int
	}{
		{"aeiou", 5},
		{"aeioa", 0},
		{"aeiaaioaaaaeiiiiouuuooaauuaeiu", 13},
		{"a", 0},
		{"aeiaeiou", 5},
		{"aeiaeioua", 5},
		{"aeaeaeiou", 5},
		{"aeiouaa", 5},
		{"faeiouaa", 5},
		{"ffaaeiouuaa", 7},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := longestBeautifulSubstring(tt.word); got != tt.want {
				t.Errorf("longestBeautifulSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
