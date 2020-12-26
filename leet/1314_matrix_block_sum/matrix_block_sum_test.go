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
			if got := matrixBlockSum2(tt.mat, tt.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixBlockSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_helper(t *testing.T) {

	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "1",
			matrix: matrix.New(3, 3),
			want: [][]int{
				{1, 3, 6},
				{5, 12, 21},
				{12, 27, 45},
			},
		},
		{
			name: "2",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: [][]int{
				{1, 3, 6, 10},
				{6, 14, 24, 36},
				{15, 33, 54, 78},
				{28, 60, 96, 136},
			},
		},
		{
			name: "3",
			matrix: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			want: [][]int{
				{1, 2, 3, 4},
				{2, 4, 6, 8},
				{3, 6, 9, 12},
				{4, 8, 12, 16},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helper(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Errorf("matrixBlockSum() = %v, want %v", tt.matrix, tt.want)
			}
		})
	}
}

func Benchmark_matrixBlockSum3(b *testing.B) {
	var res [][]int
	input := matrix.New(100, 100)
	for i := 0; i <= b.N; i++ {
		res = matrixBlockSum3(input, 5)
	}
	_ = res
}

func Benchmark_matrixBlockSum2(b *testing.B) {
	var res [][]int
	input := matrix.New(100, 100)
	for i := 0; i <= b.N; i++ {
		res = matrixBlockSum2(input, 5)
	}
	_ = res
}

func Benchmark_matrixBlockSum(b *testing.B) {
	var res [][]int
	input := matrix.New(100, 100)
	for i := 0; i <= b.N; i++ {
		res = matrixBlockSum(input, 5)
	}
	_ = res
}
