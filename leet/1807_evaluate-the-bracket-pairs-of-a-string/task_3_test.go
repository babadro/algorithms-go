package _807_evaluate_the_bracket_pairs_of_a_string

import "testing"

func Test_evaluate(t *testing.T) {

	tests := []struct {
		s         string
		knowledge [][]string
		want      string
	}{
		{
			s:         "(name)is(age)yearsold",
			knowledge: [][]string{{"name", "bob"}, {"age", "two"}},
			want:      "bobistwoyearsold",
		},
		{
			s:         "hi(name)",
			knowledge: [][]string{{"a", "b"}},
			want:      "hi?",
		},
		{
			s:         "h",
			knowledge: [][]string{},
			want:      "h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := evaluate(tt.s, tt.knowledge); got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
