package _1601_maximum_number_of_achievable_transfer_requests

import "testing"

func Test_maximumRequests(t *testing.T) {
	tests := []struct {
		n        int
		requests [][]int
		want     int
	}{
		{
			5, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}, 5,
		},
		{3, [][]int{{0, 0}, {1, 2}, {2, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maximumRequests(tt.n, tt.requests); got != tt.want {
				t.Errorf("maximumRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
