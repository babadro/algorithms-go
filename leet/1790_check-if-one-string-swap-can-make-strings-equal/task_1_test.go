package _1790_check_if_one_string_swap_can_make_strings_equal

import "testing"

func Test_areAlmostEqual(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			"1",
			"bank",
			"kanb",
			true,
		},
		{
			"2",
			"kelb",
			"kelb",
			true,
		},
		{
			"3",
			"bank",
			"dank",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areAlmostEqual(tt.s1, tt.s2); got != tt.want {
				t.Errorf("areAlmostEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
