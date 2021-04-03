package _537_complex_number_multiplication

import "testing"

func Test_complexNumberMultiply(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want string
	}{
		{"1+1i", "1+1i", "0+2i"},
		{"1+-1i", "1+-1i", "0+-2i"},
		{"0+0i", "0+0i", "0+0i"},
		{"0-0i", "0-0i", "0+0i"},
		{"3+-4i", "5+-6i", "-9+-38i"},
		{"3+4i", "5+-6i", "39+2i"},
		{"3+-4i", "5+6i", "39+-2i"},
		{"3+4i", "5+6i", "-9+38i"},
	}
	for _, tt := range tests {
		t.Run(tt.a+"_"+tt.b, func(t *testing.T) {
			if got := complexNumberMultiply(tt.a, tt.b); got != tt.want {
				t.Errorf("complexNumberMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
