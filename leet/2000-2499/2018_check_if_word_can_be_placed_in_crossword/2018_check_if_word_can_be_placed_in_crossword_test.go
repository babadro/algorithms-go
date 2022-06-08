package _2018_check_if_word_can_be_placed_in_crossword

import (
	"fmt"
	"testing"
)

func Test_placeWordInCrossword(t *testing.T) {
	tests := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{
			[][]byte{
				{'#', ' ', '#'},
				{' ', ' ', '#'},
				{'#', 'c', ' '},
			},
			"abc", true,
		},
		{
			[][]byte{
				{' ', '#', 'a'},
				{' ', '#', 'c'},
				{' ', '#', 'a'},
			}, "ac", false,
		},
		{
			[][]byte{
				{'#', ' ', '#'},
				{'#', ' ', '#'},
				{'#', ' ', 'c'},
			}, "ca", true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.board), func(t *testing.T) {
			if got := placeWordInCrossword(tt.board, tt.word); got != tt.want {
				t.Errorf("placeWordInCrossword() = %v, want %v", got, tt.want)
			}
		})
	}
}
