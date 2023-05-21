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

func Test_ParseArrFromFile(t *testing.T) {
	tests := []struct {
		filePath string
		want     interface{}
		wantType string
	}{
		{"testdata/string_arr.txt", []*string{nil, Ptr(""), Ptr("a"), Ptr("ab"), Ptr("abc")}, "string"},
		{"testdata/2d_int_arr.txt", []*[]int{nil, Ptr([]int{}), Ptr([]int{1}), Ptr([]int{1, 2})}, "2d_int_arr"},
	}
	for _, tt := range tests {
		t.Run(tt.filePath, func(t *testing.T) {
			var got interface{}
			switch tt.wantType {
			case "string":
				got = ParseArrFromFile[string](tt.filePath)
			case "2d_int_arr":
				got = ParseArrFromFile[[]int](tt.filePath)
			default:
				t.Errorf("type %s not implemented", tt.wantType)
			}

			require.Equal(t, tt.want, got)
		})
	}
}

func Test_ParseArr(t *testing.T) {
	tests := []struct {
		val      string
		want     interface{}
		wantType string
	}{
		{`[null, "", "a","ab","abc"]`, []*string{nil, Ptr(""), Ptr("a"), Ptr("ab"), Ptr("abc")}, "string"},
		{`[null, [], [1], [1,2]]`, []*[]int{nil, Ptr([]int{}), Ptr([]int{1}), Ptr([]int{1, 2})}, "2d_int_arr"},
	}
	for _, tt := range tests {
		t.Run(tt.val, func(t *testing.T) {
			var got interface{}
			switch tt.wantType {
			case "string":
				got = ParseArr[string](tt.val)
			case "2d_int_arr":
				got = ParseArr[[]int](tt.val)
			default:
				t.Errorf("type %s not implemented", tt.wantType)
			}

			require.Equal(t, tt.want, got)
		})
	}
}
