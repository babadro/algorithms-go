package _1386_cinema_seat_allocation

import "testing"

func Test_maxNumberOfFamilies(t *testing.T) {

	tests := []struct {
		reservedSeats [][]int
		n             int
		want          int
	}{
		{reservedSeats: [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}}, n: 3, want: 4},
		{reservedSeats: [][]int{{2, 1}, {1, 8}, {2, 6}}, n: 2, want: 2},
		{reservedSeats: [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}}, n: 4, want: 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxNumberOfFamilies(tt.n, tt.reservedSeats); got != tt.want {
				t.Errorf("maxNumberOfFamilies() = %v, want %v", got, tt.want)
			}
		})
	}
}
