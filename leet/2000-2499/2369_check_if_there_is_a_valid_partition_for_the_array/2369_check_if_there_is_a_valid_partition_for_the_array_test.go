package _2369_check_if_there_is_a_valid_partition_for_the_array

import "testing"

func Test_validPartition(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{4, 4, 4, 5, 6}, true},
		{[]int{1, 1, 1, 2}, false},
		{[]int{993335, 993336, 993337, 993338, 993339, 993340, 993341}, false},
		{tle, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := validPartition(tt.nums); got != tt.want {
				t.Errorf("validPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
