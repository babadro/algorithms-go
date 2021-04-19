package _784_check_if_binary_string_has_at_most_one_segment_of_ones

import "testing"

func Test_checkOnesSegment(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{s: "1001", want: false},
		{s: "110", want: true},
		{s: "1", want: true},
		{s: "0", want: true},
		{s: "01", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkOnesSegment(tt.s); got != tt.want {
				t.Errorf("checkOnesSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}
