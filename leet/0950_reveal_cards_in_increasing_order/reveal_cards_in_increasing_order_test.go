package _950_reveal_cards_in_increasing_order

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_deckRevealedIncreasing(t *testing.T) {
	tests := []struct {
		deck []int
		want []int
	}{
		{[]int{17, 13, 11, 2, 3, 5, 7}, []int{2, 13, 3, 11, 5, 17, 7}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.deck), func(t *testing.T) {
			if got := deckRevealedIncreasing(tt.deck); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deckRevealedIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
