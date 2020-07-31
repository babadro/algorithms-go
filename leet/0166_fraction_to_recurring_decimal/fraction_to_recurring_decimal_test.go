package _166_fraction_to_recurring_decimal

import (
	"testing"
)

// TODO 1 1/6 = 0.1(6) fails
func TestFractionToDecimal(t *testing.T) {
	t.Log(fractionToDecimal(2, 1))
}
