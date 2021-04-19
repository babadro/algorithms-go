package _289_game_of_life

import (
	"reflect"
	"testing"
)

func Test_gameOfLife(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  [][]int
	}{
		{
			name: "1",
			board: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		{
			name:  "2",
			board: [][]int{},
			want:  [][]int{},
		},
		{
			name: "3",
			board: [][]int{
				{0, 1, 0},
			},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "4",
			board: [][]int{
				{1, 1, 1},
			},
			want: [][]int{
				{0, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameOfLife(tt.board)
			if !reflect.DeepEqual(tt.board, tt.want) {
				t.Errorf("got = %v, want %v", tt.board, tt.want)
			}
		})
	}
}
