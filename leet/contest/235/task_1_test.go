package _35

import "testing"

func Test_truncateSentence(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		s    string
		k    int
		want string
	}{
		{
			s:    "Hello how are you Contestant",
			k:    4,
			want: "Hello how are you",
		},
		{
			s:    "What is the solution to this problem",
			k:    4,
			want: "What is the solution",
		},
		{
			s:    "single",
			k:    1,
			want: "single",
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := truncateSentence(tt.s, tt.k); got != tt.want {
				t.Errorf("truncateSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
