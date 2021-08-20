package subset_sum

import "testing"

func Test_canPartition(t *testing.T) {
	tests := []struct {
		nums []int
		sum  int
		want bool
	}{
		{[]int{1, 2, 3, 4}, 10, true},
		{[]int{14, 9, 8, 4, 3, 2}, 20, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := canPartition(tt.nums, tt.sum); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
