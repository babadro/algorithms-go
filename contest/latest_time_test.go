package contest

import "testing"

func Test_maximumTime(t *testing.T) {
	tests := []struct {
		time string
		want string
	}{
		{"2?:?0", "23:50"},
		{"0?:3?", "09:39"},
		{"1?:22", "19:22"},
		{"??:??", "23:59"},
		{"?4:03", "14:03"},
	}
	for _, tt := range tests {
		t.Run(tt.time, func(t *testing.T) {
			if got := maximumTime(tt.time); got != tt.want {
				t.Errorf("maximumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
