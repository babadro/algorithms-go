package _1314_matrix_block_sum

import (
	"github.com/babadro/algorithms-go/matrix"
	"reflect"
	"testing"
)

func Test_matrixBlockSum(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		K    int
		want [][]int
	}{
		{
			name: "1",
			mat:  matrix.New(3, 3),
			K:    1,
			want: [][]int{
				{12, 21, 16},
				{27, 45, 33},
				{24, 39, 28},
			},
		},
		{
			name: "2",
			mat:  matrix.New(3, 3),
			K:    2,
			want: [][]int{
				{45, 45, 45},
				{45, 45, 45},
				{45, 45, 45},
			},
		},
		{
			name: "3",
			mat: [][]int{
				{1, 2, 3},
			},
			K: 1,
			want: [][]int{
				{3, 6, 5},
			},
		},
		{
			name: "4",
			mat: [][]int{
				{1},
				{2},
				{3},
			},
			K: 1,
			want: [][]int{
				{3},
				{6},
				{5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixBlockSum(tt.mat, tt.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixBlockSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
