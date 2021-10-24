package __maximum_ribbon_cut

import "testing"

func Test_maximumRibbonCutTopDown(t *testing.T) {
	tests := []struct {
		lengths []int
		length  int
		want    int
	}{
		{[]int{2, 3, 5}, 5, 2},
		{[]int{2, 3}, 7, 3},
		{[]int{3, 5, 7}, 13, 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maximumRibbonBottomUp(tt.lengths, tt.length); got != tt.want {
				t.Errorf("maximumRibbonCut() = %v, want %v", got, tt.want)
			}
		})
	}
}
