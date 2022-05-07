package _3_next_letter

import "testing"

func Test_searchNextLetter(t *testing.T) {
	tests := []struct {
		letters string
		key     byte
		want    byte
	}{
		{"acfh", 'f', 'h'},
		{"acfh", 'b', 'c'},
		{"acfh", 'm', 'a'},
		{"acfh", 'h', 'a'},
		{"abbbbcf", 'b', 'c'},
	}
	for _, tt := range tests {
		t.Run(tt.letters+"_"+string(tt.key), func(t *testing.T) {
			if got := searchNextLetter(tt.letters, tt.key); got != tt.want {
				t.Errorf("searchNextLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
