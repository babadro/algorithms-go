package _1281_substract_the_product_and_sum_digits_integer

import "testing"

func Test_subtractProductAndSum(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 234, 15},
		{"2", 4421, 21},
		{"3", 111, -2},
		{"4", 0, 0},
		{"5", 1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtractProductAndSum(tt.n); got != tt.want {
				t.Errorf("subtractProductAndSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
