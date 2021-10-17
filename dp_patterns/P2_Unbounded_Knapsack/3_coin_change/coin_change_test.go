package __coin_change

import "testing"

func TestCoinChangeBruteForce(t *testing.T) {
	tests := []struct {
		denominations []int
		total         int
		want          int
	}{
		{[]int{1, 2, 3}, 5, 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CoinChangeBruteForce(tt.denominations, tt.total); got != tt.want {
				t.Errorf("CoinChangeBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
