package _502_ipo

import "testing"

func Test_findMaximizedCapital(t *testing.T) {
	tests := []struct {
		k       int
		w       int
		profits []int
		capital []int
		want    int
	}{
		{2, 0, []int{1, 2, 3}, []int{0, 1, 1}, 4},
		{3, 0, []int{1, 2, 3}, []int{0, 1, 2}, 6},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findMaximizedCapital(tt.k, tt.w, tt.profits, tt.capital); got != tt.want {
				t.Errorf("findMaximizedCapital() = %v, want %v", got, tt.want)
			}
		})
	}
}
