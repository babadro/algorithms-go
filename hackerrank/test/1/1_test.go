package _1

import "testing"

func Test_bioHazard(t *testing.T) {
	tests := []struct {
		n         int32
		allergic  []int32
		poisonous []int32
		want      int64
	}{
		//{3, []int32{2, 1, 3}, []int32{3, 3, 1}, 4},
		//{4, []int32{1, 2}, []int32{3, 4}, 7},
		//{5, []int32{1, 2}, []int32{3, 5}, 11},
		//{5, []int32{1, 2}, []int32{3, 5}, 11},
		{8, []int32{2, 3, 4, 3}, []int32{8, 5, 6, 4}, 18},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := bioHazard(tt.n, tt.allergic, tt.poisonous); got != tt.want {
				t.Errorf("bioHazard() = %v, want %v", got, tt.want)
			}
		})
	}
}
