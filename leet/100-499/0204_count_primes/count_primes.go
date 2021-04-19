package _204_count_primes

import "math"

// passed. tptl. need to understand
func countPrimes(n int) int {
	var counter int
	arr := make([]bool, n)
	for i := 2; i < n; i++ {
		if arr[i] {
			continue
		}
		counter++
		for j := i * i; j < n; j += i {
			arr[j] = true
		}
	}
	return counter
}

func countPrimes1(n int) int {
	if n <= 2 {
		return 0
	}
	arr := make([]bool, n)
	i, counter := 1, 0
	for {
		i++
		if arr[i] || !isPrime(i) {
			continue
		}
		num := i * i
		if num >= n {
			break
		}
		for num < n {
			if !arr[num] {
				arr[num] = true
				counter++
			}
			num = num + i
		}
	}
	return n - counter - 2
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
