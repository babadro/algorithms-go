package __find_the_corrupt_pair

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findCorruptNums(t *testing.T) {
	tests := []struct {
		nums  []int
		want  int
		want1 int
	}{
		{[]int{3, 1, 2, 5, 2}, 2, 4},
		{[]int{3, 1, 2, 3, 6, 4}, 3, 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, got1 := findCorruptNums(tt.nums)
			if got > got1 {
				got, got1 = got1, got
			}
			want, want1 := tt.want, tt.want1
			if want > want1 {
				want, want1 = want1, want
			}

			require.Equal(t, want, got)
			require.Equal(t, want1, got1)
		})
	}
}
