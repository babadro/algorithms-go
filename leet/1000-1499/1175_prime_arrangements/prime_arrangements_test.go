package _1175_prime_arrangements

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_isPrime(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{9, false},
		{29, true},
		{53, true},
		{97, true},
		{99, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.num), func(t *testing.T) {
			if got := isPrime(tt.num); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

// for generating prefix sum:
// var primeSums = []int{0, 0, 1, 2, 2, 3, 3, 4, 4, 4, 4, 5, 5, 6, 6, 6, 6, 7, 7, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 10, 10, 11, 11, 11, 11, 11, 11, 12, 12, 12, 12, 13, 13, 14, 14, 14, 14, 15, 15, 15, 15, 15, 15, 16, 16, 16, 16, 16, 16, 17, 17, 18, 18, 18, 18, 18, 18, 19, 19, 19, 19, 20, 20, 21, 21, 21, 21, 21, 21, 22, 22, 22, 22, 23, 23, 23, 23, 23, 23, 24, 24, 24, 24, 24, 24, 24, 24, 25, 25, 25, 25}
func Test_isPrime2(t *testing.T) {
	s := "[]int{0"
	sum := 0
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			sum++
		}

		s += "," + strconv.Itoa(sum)
	}

	s += "}"

	t.Log(s)
}

func Test_generate_answers(t *testing.T) {
	res := "[]int{0"
	for i := 1; i <= 100; i++ {
		res += "," + strconv.Itoa(numPrimeArrangements(i))
	}

	res += "}"

	t.Log(res)
}

func Test_factorialDivider(t *testing.T) {
	for i := 1; i <= 15; i++ {
		t.Log(factorialDivider(i))
	}
}

func Test_numPrimeArrangements(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 4},
		{5, 12},
		{100, 682289015},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			if got := numPrimeArrangements2(tt.n); got != tt.want {
				t.Errorf("numPrimeArrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}
