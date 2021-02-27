package _1175_prime_arrangements

var primeSums = []int{0, 0, 1, 2, 2, 3, 3, 4, 4, 4, 4, 5, 5, 6, 6, 6, 6, 7, 7, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 10, 10, 11, 11, 11, 11, 11, 11, 12, 12, 12, 12, 13, 13, 14, 14, 14, 14, 15, 15, 15, 15, 15, 15, 16, 16, 16, 16, 16, 16, 17, 17, 18, 18, 18, 18, 18, 18, 19, 19, 19, 19, 20, 20, 21, 21, 21, 21, 21, 21, 22, 22, 22, 22, 23, 23, 23, 23, 23, 23, 24, 24, 24, 24, 24, 24, 24, 24, 25, 25, 25, 25}

// passed. tptl
func numPrimeArrangements(n int) int {
	primeCount := primeSums[n]

	return (factorialDivider(n-primeCount) * factorialDivider(primeCount)) % 1_000_000_007
}

func factorialDivider(num int) int {
	res := 1
	for i := 1; i <= num; i++ {
		res = (res * i) % 1_000_000_007
	}

	return res
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
