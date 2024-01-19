package _247_strobogrammatic_number_2

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findStrobogrammatic(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{2, []string{"11", "69", "88", "96"}},
		{3, []string{"101", "111", "181", "609", "619", "689", "808", "818", "888", "906", "916", "986"}},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			got := findStrobogrammatic(tt.n)

			sort.Strings(got)
			sort.Strings(tt.want)

			require.Equal(t, tt.want, got)
		})
	}
}
