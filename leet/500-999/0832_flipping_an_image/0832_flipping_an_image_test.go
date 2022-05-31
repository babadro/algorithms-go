package _832_flipping_an_image

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	tests := []struct {
		image [][]int
		want  [][]int
	}{
		{
			[][]int{
				{1, 1, 0},
				{1, 0, 1},
				{0, 0, 0},
			},
			[][]int{
				{1, 0, 0}, {0, 1, 0}, {1, 1, 1},
			},
		},
		{
			[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			[][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.image), func(t *testing.T) {
			if got := flipAndInvertImage(tt.image); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
