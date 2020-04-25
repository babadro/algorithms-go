package _204_count_primes

import "math"

// TODO
func countPrimes(n int) int {

	arr := make([]bool, n)
	for {
		break
	}
	counter := 0
	for _, notPrime := range arr {
		if !notPrime {
			counter++
		}
	}
	return counter
}

func countPrimesNaive(n int) int {
	counter := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			counter++
		}
	}
	return counter
}

func isPrime(n int) bool {
	to := int(math.Sqrt(float64(n)))
	for i := 2; i <= to; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
