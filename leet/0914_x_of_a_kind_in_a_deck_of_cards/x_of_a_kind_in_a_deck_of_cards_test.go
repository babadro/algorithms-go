package _914_x_of_a_kind_in_a_deck_of_cards

import "testing"

func Test_hasGroupsSizeX(t *testing.T) {

	tests := []struct {
		name string
		deck []int
		want bool
	}{
		{"1", []int{1, 2, 3, 4, 4, 3, 2, 1}, true},
		//{"2", []int{1,1,1,2,2,2,3,3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasGroupsSizeX(tt.deck); got != tt.want {
				t.Errorf("hasGroupsSizeX() = %v, want %v", got, tt.want)
			}
		})
	}
}
