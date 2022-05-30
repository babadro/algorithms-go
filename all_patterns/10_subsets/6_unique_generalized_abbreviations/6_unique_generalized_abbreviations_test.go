package __unique_generalized_abbreviations

import (
	"reflect"
	"sort"
	"testing"
)

func Test_generateGeneralizedAbbreviation(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{"BAT", []string{"BAT", "BA1", "B1T", "B2", "1AT", "1A1", "2T", "3"}},
		{"code", []string{"code", "cod1", "co1e", "co2", "c1de", "c1d1", "c2e", "c3", "1ode", "1od1", "1o1e", "1o2",
			"2de", "2d1", "3e", "4"}},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := generateGeneralizedAbbreviation(tt.s)
			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateGeneralizedAbbreviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
