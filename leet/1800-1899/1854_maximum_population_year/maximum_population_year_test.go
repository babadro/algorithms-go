package _1854_maximum_population_year

import (
	"fmt"
	"testing"
)

func Test_maximumPopulation(t *testing.T) {
	tests := []struct {
		logs [][]int
		want int
	}{
		{[][]int{{1993, 1999}, {2000, 2010}}, 1993},
		{[][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}, 1960},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.logs), func(t *testing.T) {
			if got := maximumPopulation2(tt.logs); got != tt.want {
				t.Errorf("maximumPopulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
