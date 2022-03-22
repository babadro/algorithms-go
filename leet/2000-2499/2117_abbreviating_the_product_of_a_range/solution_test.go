package _2117_abbreviating_the_product_of_a_range

import (
	"fmt"
	"testing"
)

func Test_abbreviateProduct(t *testing.T) {
	tests := []struct {
		left  int
		right int
		want  string
	}{
		{1, 4, "24e0"},
		{2, 11, "399168e2"},
		{371, 375, "7219856259e3"},
		{900553, 900563, "31595...16864e2"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("left: %d, right: %d", tt.left, tt.right), func(t *testing.T) {
			if got := abbreviateProduct(tt.left, tt.right); got != tt.want {
				t.Errorf("abbreviateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
