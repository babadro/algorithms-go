package _2136_earliest_possible_day_full_bloom

import "testing"

func Test_earliestFullBloom(t *testing.T) {
	tests := []struct {
		plantTime, growTime []int
		want                int
	}{
		{[]int{1, 4, 3}, []int{2, 3, 1}, 9},
		{[]int{1, 2, 3, 2}, []int{2, 1, 2, 1}, 9},
		{[]int{1}, []int{1}, 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := earliestFullBloom(tt.plantTime, tt.growTime); got != tt.want {
				t.Errorf("earliestFullBloom() = %v, want %v", got, tt.want)
			}
		})
	}
}
