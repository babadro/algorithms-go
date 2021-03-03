package _819_most_common_word

import "testing"

func Test_mostCommonWord(t *testing.T) {
	tests := []struct {
		paragraph string
		banned    []string
		want      string
	}{
		{"Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}, "ball"},
		{"bob mum bob.", []string{"hit"}, "bob"},
	}
	for _, tt := range tests {
		t.Run(tt.paragraph, func(t *testing.T) {
			if got := mostCommonWord(tt.paragraph, tt.banned); got != tt.want {
				t.Errorf("mostCommonWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
