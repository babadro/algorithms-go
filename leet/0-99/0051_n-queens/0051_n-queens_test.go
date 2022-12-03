package _051_n_queens

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	tests := []struct {
		n    int
		want [][]string
	}{
		{
			n: 4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			n:    1,
			want: [][]string{{"Q"}},
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			if got := solveNQueens(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
