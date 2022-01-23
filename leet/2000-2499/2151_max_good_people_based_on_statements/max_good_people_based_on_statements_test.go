package _2151_max_good_people_based_on_statements

import "testing"

func Test_maximumGood(t *testing.T) {
	tests := []struct {
		statements [][]int
		want       int
	}{
		{[][]int{{2, 1, 2}, {1, 2, 2}, {2, 0, 0}}, 2},
		{[][]int{{2, 0}, {0, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maximumGood(tt.statements); got != tt.want {
				t.Errorf("maximumGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
