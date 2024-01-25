package _418_sentence_screen_fitting

import "testing"

func Test_wordsTyping(t *testing.T) {
	tests := []struct {
		sentence []string
		rows     int
		cols     int
		want     int
	}{
		{
			sentence: []string{"hello", "world"}, rows: 2, cols: 8, want: 1,
		},
		{
			[]string{"a", "bcd", "e"}, 3, 6, 2,
		},
		{
			[]string{"i", "had", "apple", "pie"}, 4, 5, 1,
		},
		{
			[]string{"f", "p", "a"}, 8, 7, 10,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := wordsTyping(tt.sentence, tt.rows, tt.cols); got != tt.want {
				t.Errorf("wordsTyping() = %v, want %v", got, tt.want)
			}
		})
	}
}
