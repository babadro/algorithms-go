package _914_x_of_a_kind_in_a_deck_of_cards

import "testing"

func Test_hasGgoroupsSizeX(t *testing.T) {

	tests := []struct {
		name string
		deck []int
		want bool
	}{
		{"1", []int{1, 2, 3, 4, 4, 3, 2, 1}, true},
		{"2", []int{1, 1, 1, 2, 2, 2, 3, 3}, false},
		{"3", []int{1}, false},
		{"4", []int{1, 1}, true},
		{"5", []int{1, 1, 2, 2, 2, 2}, true},
		{"5", []int{1, 1, 2, 3, 3, 4}, false},
		{"6", []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2}, true}, // failed.
		{"7", []int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3}, false},
		{"8", []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 5, 5, 5, 5, 6, 6, 7, 7, 8, 8}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasGroupsSizeX(tt.deck); got != tt.want {
				t.Errorf("hasGroupsSizeX() = %v, want %v", got, tt.want)
			}
		})
	}
}

//14, 6, 6, 12, 2, 4, 2, 2, 2
