package _201_Bitwise_And_of_Numbers_Range

import (
	"fmt"
	"testing"
)

func Test_rangeBitwiseAnd(t *testing.T) {
	tests := []struct {
		m    int
		n    int
		want int
	}{
		{4, 7, -1},
		{0, 0, -1},
		{1, 2147483647, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d", tt.m, tt.n), func(t *testing.T) {
			var want int
			if tt.want != -1 {
				want = tt.want
			} else {
				want = rangeBitwiseAndNaive(tt.m, tt.n)
			}

			if got := rangeBitwiseAnd(tt.m, tt.n); got != want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, want)
			}
		})
	}
}
