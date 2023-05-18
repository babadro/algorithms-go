package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseNullableNums(t *testing.T) {
	tests := []struct {
		filePath string
		want     []interface{}
	}{
		{"testdata/interfaces_arr.txt", []interface{}{1, nil, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.filePath, func(t *testing.T) {
			got := ParseNullableNums[int](tt.filePath)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_ParseNums(t *testing.T) {
	tests := []struct {
		filePath string
		want     []int
	}{
		{"testdata/ints_arr.txt", []int{1, 0, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.filePath, func(t *testing.T) {
			got := ParseNums[int](tt.filePath)
			require.Equal(t, tt.want, got)
		})
	}
}
