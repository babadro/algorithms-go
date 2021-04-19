package _1832_check_if_the_sentence_is_pangram

import "testing"

func Test_checkIfPangram(t *testing.T) {
	tests := []struct {
		sentence string
		want     bool
	}{
		{"thequickbrownfoxjumpsoverthelazydog", true},
		{"leetcode", false},
	}
	for _, tt := range tests {
		t.Run(tt.sentence, func(t *testing.T) {
			if got := checkIfPangram(tt.sentence); got != tt.want {
				t.Errorf("checkIfPangram() = %v, want %v", got, tt.want)
			}
		})
	}
}
