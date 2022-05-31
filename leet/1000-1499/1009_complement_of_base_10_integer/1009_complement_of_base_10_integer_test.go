package _1009_complement_of_base_10_integer

import (
	"strconv"
	"testing"
)

func Test_bitwiseComplement(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{5, 2},
		{7, 0},
		{10, 5},
		{0, 1},
	}
	for _, tt := range tests {
		t.Run(strconv.FormatInt(int64(tt.n), 2), func(t *testing.T) {
			if got := bitwiseComplement(tt.n); got != tt.want {
				t.Errorf("bitwiseComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}
