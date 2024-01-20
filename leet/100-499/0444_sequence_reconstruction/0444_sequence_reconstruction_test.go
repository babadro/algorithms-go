package _444_sequence_reconstruction

import "testing"

func Test_sequenceReconstruction(t *testing.T) {
	tests := []struct {
		nums      []int
		sequences [][]int
		want      bool
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}}, false},
		{[]int{1, 2, 3}, [][]int{{1, 2}}, false},
		{[]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {2, 3}}, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := sequenceReconstruction(tt.nums, tt.sequences); got != tt.want {
				t.Errorf("sequenceReconstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
