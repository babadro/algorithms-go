package __Rod_Cutting

import "testing"

func TestRodCuttingBruteForce(t *testing.T) {
	tests := []struct {
		lengths []int
		prices  []int
		n       int
		want    int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 6, 7, 10, 13}, 5, 14},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := RodCuttingBottomUp(tt.lengths, tt.prices, tt.n); got != tt.want {
				t.Errorf("RodCuttingBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
