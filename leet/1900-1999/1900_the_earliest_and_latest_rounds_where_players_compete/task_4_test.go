package _1900_the_earliest_and_latest_rounds_where_players_compete

import (
	"reflect"
	"testing"
)

func Test_earliestAndLatest(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name         string
		n            int
		firstPlayer  int
		secondPlayer int
		want         []int
	}{
		{"", 11, 2, 4, []int{3, 4}},
		{"", 5, 1, 5, []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := earliestAndLatest(tt.n, tt.firstPlayer, tt.secondPlayer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("earliestAndLatest() = %v, want %v", got, tt.want)
			}
		})
	}
}
